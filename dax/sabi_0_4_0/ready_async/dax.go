package ready_async

import (
	"reflect"
	"sync"
)

type /* error reasons */ (
	// FailToStartUpGlobalDaxSrcs is an error reason which indicates that some
	// dax sources failed to start up.
	// The field: Errors is a map of which keys are the registered names of
	// DaxSrc(s) which failed to start up, and of which values are Err having
	// their error reasons.
	FailToStartUpGlobalDaxSrcs struct {
		Errors map[string]Err
	}

	// DaxSrcIsNotFound is an error reason which indicates that a specified data
	// source is not found.
	// The field: Name is a registered name of a DaxSrc not found.
	DaxSrcIsNotFound struct {
		Name string
	}

	// FailToCreateDaxConn is an error reason which indicates that it's failed to
	// create a new connection to a data store.
	// The field: Name is a registered name of a DaxSrc which failed to create a
	// DaxConn.
	FailToCreateDaxConn struct {
		Name string
	}

	// FailToCastConn is an error reason which indicates that it is failed to
	// cast type of a DaxConn.
	// The field: Name is a registered name of a DaxSrc which created the target
	// DaxConn.
	// And the fields: FromType and ToType are the types of the source DaxConn
	// and the destination DaxConn.
	FailToCastDaxConn struct {
		Name, FromType, ToType string
	}

	// FailToCommitDaxConn is an error reason which indicates that some
	// connections failed to commit.
	// The field: Errors is a map of which keys are the registered names of
	// DaxConn which failed to commit, and of which values are Err having their
	// error reasons.
	FailToCommitDaxConn struct {
		Errors map[string]Err
	}
)

// DaxConn is an interface which represents a connection to a data store, and
// defines methods: Commit, Rollback and Close to work in a tranaction process.
type DaxConn interface {
	Commit(wg *sync.WaitGroup) Err
	Committed() Err
	Rollback(wg *sync.WaitGroup)
	Close()
}

// DaxSrc is an interface which represents a data connection source for a data
// store like database, etc., and creates a DaxConn which is a connection to a
// data store.
// This interface defines a method: CreateDaxConn to creates a DaxConn instance
// and returns its pointer.
// This interface also defines methods: SetUp and End, which makes available
// and frees this dax source.
type DaxSrc interface {
	CreateDaxConn() (DaxConn, Err)
	SetUp(wg *sync.WaitGroup) Err
	Ready() Err
	End()
}

// Dax is an interface for a set of data access methods.
// This method gets a DaxConn which is a connection to a data store by
// specified name.
// If a DaxConn is found, this method returns it, but not found, creates a new
// one with a local or global DaxSrc associated with same name.
// If there are both local and global DaxSrc with same name, the local DaxSrc
// is used.
type Dax interface {
	getDaxConn(name string) (DaxConn, Err)
}

var (
	isGlobalDaxSrcsFixed bool              = false
	globalDaxSrcMap      map[string]DaxSrc = make(map[string]DaxSrc)
)

// AddGlobalDaxSrc registers a global DaxSrc with its name to make enable to
// use DaxSrc in all transactions.
// This method ignores to add a global DaxSrc when its name is already
// registered.
// In addition, this method ignores to add any more global DaxSrc(s) after
// calling FixGlobalDaxSrcs function.
func AddGlobalDaxSrc(name string, ds DaxSrc) {
	if !isGlobalDaxSrcsFixed {
		_, exists := globalDaxSrcMap[name]
		if !exists {
			globalDaxSrcMap[name] = ds
		}
	}
}

// StartUpGlobalDaxSrcs is a function to forbid adding more global dax sources
// and to make available the registered global dax sources by calling Setup
// method of each DaxSrc.
// If even one DaxSrc fail to execute its SstUp method, this function
// executes Free methods of all global DaxSrc(s) and returns sabi.Err.
func StartUpGlobalDaxSrcs() Err {
	isGlobalDaxSrcsFixed = true

	errs := make(map[string]Err)

	var wg sync.WaitGroup

	for name, ds := range globalDaxSrcMap {
		err := ds.SetUp(&wg)
		if err.IsNotOk() {
			errs[name] = err
		}
	}

	wg.Wait()

	for name, ds := range globalDaxSrcMap {
		err := ds.Ready()
		if err.IsNotOk() {
			errs[name] = err
		}
	}

	if len(errs) > 0 {
		ShutdownGlobalDaxSrcs()
		return NewErr(FailToStartUpGlobalDaxSrcs{Errors: errs})
	}

	return Ok()
}

// ShutdownGlobalDaxSrcs is a function to terminate all global dax sources.
func ShutdownGlobalDaxSrcs() {
	for _, ds := range globalDaxSrcMap {
		ds.End()
	}
}

// DaxBase is an interface which works as a front of an implementation as a
// base of data connection sources, and defines methods: SetUpLocalDaxSrc and
// FreeLocalDaxSrc.
//
// SetUpLocalDaxSrc method registered a DaxSrc with a name in this
// implementation, but  ignores to add a local DaxSrc when its name is already
// registered.
// In addition, this method ignores to add local DaxSrc(s) while the
// transaction is processing.
//
// This interface inherits Dax interface to get a DaxConn by a name.
// Also, this has unexported methods for a transaction process.
type DaxBase interface {
	Dax
	SetUpLocalDaxSrc(name string, ds DaxSrc) Err
	FreeLocalDaxSrc(name string)
	FreeAllLocalDaxSrcs()
	begin()
	commit() Err
	rollback()
	end()
}

type daxBaseImpl struct {
	isLocalDaxSrcsFixed bool
	localDaxSrcMap      map[string]DaxSrc
	daxConnMap          map[string]DaxConn
	daxConnMutex        sync.Mutex
}

// NewDaxBase is a function which creates a new DaxBase.
func NewDaxBase() DaxBase {
	return &daxBaseImpl{
		isLocalDaxSrcsFixed: false,
		localDaxSrcMap:      make(map[string]DaxSrc),
		daxConnMap:          make(map[string]DaxConn),
	}
}

func (base *daxBaseImpl) SetUpLocalDaxSrc(name string, ds DaxSrc) Err {
	if !base.isLocalDaxSrcsFixed {
		_, exists := base.localDaxSrcMap[name]
		if !exists {
			var wg sync.WaitGroup

			err := ds.SetUp(&wg)
			if err.IsNotOk() {
				return err
			}

			wg.Wait()

			err = ds.Ready()
			if err.IsNotOk() {
				return err
			}

			base.localDaxSrcMap[name] = ds
		}
	}

	return Ok()
}

func (base *daxBaseImpl) FreeLocalDaxSrc(name string) {
	if !base.isLocalDaxSrcsFixed {
		ds, exists := base.localDaxSrcMap[name]
		if exists {
			delete(base.localDaxSrcMap, name)
			ds.End()
		}
	}
}

func (base *daxBaseImpl) FreeAllLocalDaxSrcs() {
	if !base.isLocalDaxSrcsFixed {
		for _, ds := range base.localDaxSrcMap {
			ds.End()
		}

		base.localDaxSrcMap = make(map[string]DaxSrc)
	}
}

func (base *daxBaseImpl) getDaxConn(name string) (DaxConn, Err) {
	conn := base.daxConnMap[name]
	if conn != nil {
		return conn, Ok()
	}

	ds := base.localDaxSrcMap[name]
	if ds == nil {
		ds = globalDaxSrcMap[name]
	}
	if ds == nil {
		return nil, NewErr(DaxSrcIsNotFound{Name: name})
	}

	base.daxConnMutex.Lock()
	defer base.daxConnMutex.Unlock()

	conn = base.daxConnMap[name]
	if conn != nil {
		return conn, Ok()
	}

	var err Err
	conn, err = ds.CreateDaxConn()
	if !err.IsOk() {
		return nil, NewErr(FailToCreateDaxConn{Name: name}, err)
	}

	base.daxConnMap[name] = conn

	return conn, Ok()
}

func (base *daxBaseImpl) begin() {
	base.isLocalDaxSrcsFixed = true
	isGlobalDaxSrcsFixed = true
}

type namedErr struct {
	name string
	err  Err
}

func (base *daxBaseImpl) commit() Err {
	errs := make(map[string]Err)

	var wg sync.WaitGroup

	for name, conn := range base.daxConnMap {
		err := conn.Commit(&wg)
		if err.IsNotOk() {
			errs[name] = err
		}
	}

	wg.Wait()

	for name, conn := range base.daxConnMap {
		err := conn.Committed()
		if err.IsNotOk() {
			errs[name] = err
		}
	}

	if len(errs) > 0 {
		return NewErr(FailToCommitDaxConn{Errors: errs})
	}

	return Ok()
}

func (base *daxBaseImpl) rollback() {
	var wg sync.WaitGroup

	for _, conn := range base.daxConnMap {
		conn.Rollback(&wg)
	}

	wg.Wait()
}

func (base *daxBaseImpl) end() {
	for _, conn := range base.daxConnMap {
		conn.Close()
	}

	base.daxConnMap = make(map[string]DaxConn)
	base.isLocalDaxSrcsFixed = false
}

// GetDaxConn is a function to cast type of DaxConn instance.
// If it's failed to cast to a destination type, this function returns an Err
// of a reason: FailToGetDaxConn.
func GetDaxConn[C DaxConn](dax Dax, name string) (C, Err) {
	conn, err := dax.getDaxConn(name)
	if err.IsOk() {
		casted, ok := conn.(C)
		if ok {
			return casted, err
		}

		var from string
		t := reflect.TypeOf(conn)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			from = "*" + t.Name() + " (" + t.PkgPath() + ")"
		} else {
			from = t.Name() + " (" + t.PkgPath() + ")"
		}

		var to string
		t = reflect.TypeOf(casted)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			to = "*" + t.Name() + " (" + t.PkgPath() + ")"
		} else {
			to = t.Name() + " (" + t.PkgPath() + ")"
		}
		err = NewErr(FailToCastDaxConn{Name: name, FromType: from, ToType: to})
	}

	return *new(C), err
}
