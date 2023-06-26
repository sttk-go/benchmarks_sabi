package orderedmap_isasync_sync

import (
	"reflect"
	"sync"

	om "github.com/sttk/orderedmap"
)

type (
	FailToStartUpGlobalDaxSrcs struct {
		Errors map[string]Err
	}

	DaxSrcIsNotFound struct {
		Name string
	}

	FailToCreateDaxConn struct {
		Name string
	}

	CreatedDaxConnIsNil struct {
		Name string
	}

	FailToCastDaxConn struct {
		Name, FromType, ToType string
	}

	FailToCommitDaxConn struct {
		Errors map[string]Err
	}
)

var (
	isGlobalDaxSrcsFixed bool                   = false
	globalDaxSrcMap      om.Map[string, DaxSrc] = om.New[string, DaxSrc]()
	globalErrors         map[string]Err         = make(map[string]Err)
)

type DaxConn interface {
	IsAsync() bool
	Commit(wg *sync.WaitGroup) Err
	Committed() Err
	Rollback(wg *sync.WaitGroup)
	Close()
}

type DaxSrc interface {
	CreateDaxConn() (DaxConn, Err)
	IsAsync() bool
	SetUp(wg *sync.WaitGroup) Err
	Ready() Err
	End()
}

func AddGlobalDaxSrc(name string, ds DaxSrc) {
	globalDaxSrcMap.LoadOrStore(name, ds)
}

func StartUpGlobalDaxSrcs() Err {
	var wg *sync.WaitGroup

	for ent := globalDaxSrcMap.Front(); ent != nil; ent = ent.Next() {
		ds := ent.Value()
		if ds.IsAsync() {
			if wg == nil {
				wg = &sync.WaitGroup{}
			}
		}
		err := ds.SetUp(wg)
		if err.IsNotOk() {
			globalErrors[ent.Key()] = err
		}
	}

	if wg != nil {
		wg.Wait()
	}

	for ent := globalDaxSrcMap.Front(); ent != nil; ent = ent.Next() {
		err := ent.Value().Ready()
		if err.IsNotOk() {
			globalErrors[ent.Key()] = err
		}
	}

	if len(globalErrors) > 0 {
		ShutdownGlobalDaxSrcs()
		return NewErr(FailToStartUpGlobalDaxSrcs{Errors: globalErrors})
	}

	return Ok()
}

func ShutdownGlobalDaxSrcs() {
	for ent := globalDaxSrcMap.Front(); ent != nil; ent = ent.Next() {
		ent.Value().End()
	}
}

type Dax interface {
	getDaxConn(name string) (DaxConn, Err)
}

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
	localDaxSrcMap      om.Map[string, DaxSrc]
	daxConnMap          om.Map[string, DaxConn]
	daxConnMutex        sync.Mutex
}

func NewDaxBase(wg *sync.WaitGroup) DaxBase {
	return &daxBaseImpl{
		isLocalDaxSrcsFixed: false,
		localDaxSrcMap:      om.New[string, DaxSrc](),
		daxConnMap:          om.New[string, DaxConn](),
	}
}

func (base *daxBaseImpl) SetUpLocalDaxSrc(name string, ds DaxSrc) Err {
	if !base.isLocalDaxSrcsFixed {
		_, loaded := base.localDaxSrcMap.LoadOrStore(name, ds)
		if loaded {
			return Ok()
		}

		err := ds.SetUp(nil)
		if err.IsNotOk() {
			return err
		}

		err = ds.Ready()
		if err.IsNotOk() {
			return err
		}
	}

	return Ok()
}

func (base *daxBaseImpl) FreeLocalDaxSrc(name string) {
	if !base.isLocalDaxSrcsFixed {
		ds, exists := base.localDaxSrcMap.LoadAndLdelete(name)
		if exists {
			ds.End()
		}
	}
}

func (base *daxBaseImpl) FreeAllLocalDaxSrcs() {
	if !base.isLocalDaxSrcsFixed {
		for ent := base.localDaxSrcMap.Front(); ent != nil; ent = ent.Next() {
			ent.Value().End()
		}

		base.localDaxSrcMap = om.New[string, DaxSrc]()
	}
}

func (base *daxBaseImpl) begin() {
	base.isLocalDaxSrcsFixed = true
	isGlobalDaxSrcsFixed = true
}

func (base *daxBaseImpl) commit() Err {
	errs := make(map[string]Err)

	var wg *sync.WaitGroup

	for ent := base.daxConnMap.Front(); ent != nil; ent = ent.Next() {
		conn := ent.Value()
		if conn.IsAsync() {
			if wg == nil {
				wg = &sync.WaitGroup{}
			}
		}
		err := conn.Commit(wg)
		if err.IsNotOk() {
			errs[ent.Key()] = err
		}
	}

	if wg == nil {
		wg.Wait()
	}

	for ent := base.daxConnMap.Front(); ent != nil; ent = ent.Next() {
		err := ent.Value().Committed()
		if err.IsNotOk() {
			errs[ent.Key()] = err
		}
	}

	if len(errs) > 0 {
		return NewErr(FailToCommitDaxConn{Errors: errs})
	}

	return Ok()
}

func (base *daxBaseImpl) rollback() {
	var wg *sync.WaitGroup

	for ent := base.daxConnMap.Front(); ent != nil; ent = ent.Next() {
		conn := ent.Value()
		if conn.IsAsync() {
			if wg == nil {
				wg = &sync.WaitGroup{}
			}
		}
		conn.Rollback(wg)
	}

	if wg != nil {
		wg.Wait()
	}
}

func (base *daxBaseImpl) end() {
	for ent := base.daxConnMap.Front(); ent != nil; ent = ent.Next() {
		ent.Value().Close()
	}

	base.daxConnMap = om.New[string, DaxConn]()
	base.isLocalDaxSrcsFixed = false
}

func (base *daxBaseImpl) getDaxConn(name string) (DaxConn, Err) {
	conn, loaded := base.daxConnMap.Load(name)
	if loaded {
		return conn, Ok()
	}

	base.daxConnMutex.Lock()
	defer base.daxConnMutex.Unlock()

	fn := func() (DaxConn, error) {
		ds, exists := base.localDaxSrcMap.Load(name)
		if !exists {
			ds, exists = globalDaxSrcMap.Load(name)
		}
		if !exists {
			return nil, NewErr(DaxSrcIsNotFound{Name: name})
		}

		conn, err := ds.CreateDaxConn()
		if err.IsNotOk() {
			return nil, NewErr(FailToCreateDaxConn{Name: name}, err)
		}
		if conn == nil {
			return nil, NewErr(CreatedDaxConnIsNil{Name: name})
		}
		return conn, nil
	}

	conn, _, e := base.daxConnMap.LoadOrStoreFunc(name, fn)
	if e != nil {
		return nil, e.(Err)
	}
	return conn, Ok()
}

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
