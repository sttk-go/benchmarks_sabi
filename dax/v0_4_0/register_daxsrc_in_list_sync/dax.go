package register_daxsrc_in_list_sync

import (
	"reflect"
	"sync"
	// om "github.com/sttk/orderedmap"
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

type daxSrcEntry struct {
	name  string
	ds    DaxSrc
	prev  *daxSrcEntry
	next  *daxSrcEntry
	local bool
}

type daxSrcList struct {
	head *daxSrcEntry
	last *daxSrcEntry
}

type daxConnEntry struct {
	name string
	conn DaxConn
	prev *daxConnEntry
	next *daxConnEntry
}

type daxConnList struct {
	head *daxConnEntry
	last *daxConnEntry
}

var (
	isGlobalDaxSrcsFixed bool       = false
	globalDaxSrcList     daxSrcList = daxSrcList{}
)

type DaxConn interface {
	Commit(wg *sync.WaitGroup) Err
	Committed() Err
	Rollback(wg *sync.WaitGroup)
	Close()
}

type DaxSrc interface {
	SetUp(wg *sync.WaitGroup) Err
	Ready() Err
	End()
	CreateDaxConn() (DaxConn, Err)
}

func AddGlobalDaxSrc(name string, ds DaxSrc) {
	ent := &daxSrcEntry{name: name, ds: ds}
	if globalDaxSrcList.head == nil {
		globalDaxSrcList.head = ent
		globalDaxSrcList.last = ent
	} else {
		ent.prev = globalDaxSrcList.last
		globalDaxSrcList.last.next = ent
		globalDaxSrcList.last = ent
	}
}

func StartUpGlobalDaxSrcs() Err {
	var wg sync.WaitGroup

	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		err := ent.ds.SetUp(&wg)
		if err.IsNotOk() {
			wg.Wait()
			ShutdownGlobalDaxSrcs()
			errs := make(map[string]Err)
			errs[ent.name] = err
			for ent := ent.prev; ent != nil; ent = ent.prev {
				err = ent.ds.Ready()
				if err.IsNotOk() {
					errs[ent.name] = err
				}
			}
			return NewErr(FailToStartUpGlobalDaxSrcs{Errors: errs})
		}
	}

	wg.Wait()

	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		err := ent.ds.Ready()
		if err.IsNotOk() {
			ShutdownGlobalDaxSrcs()
			errs := make(map[string]Err)
			errs[ent.name] = err
			for ent := ent.prev; ent != nil; ent = ent.prev {
				err = ent.ds.Ready()
				if err.IsNotOk() {
					errs[ent.name] = err
				}
			}
			return NewErr(FailToStartUpGlobalDaxSrcs{Errors: errs})
		}
	}

	return Ok()
}

func ShutdownGlobalDaxSrcs() {
	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		ent.ds.End()
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
	localDaxSrcList     daxSrcList
	daxSrcEntryMap      map[string](*daxSrcEntry)
	daxConnList         daxConnList
	daxConnMap          map[string](*daxConnEntry)
	//daxConnMap          om.Map[string, DaxConn]
	daxConnMutex sync.Mutex
}

func NewDaxBase() DaxBase {
	base := daxBaseImpl{
		isLocalDaxSrcsFixed: false,
		localDaxSrcList:     daxSrcList{},
		daxSrcEntryMap:      make(map[string]*daxSrcEntry),
		daxConnList:         daxConnList{},
		daxConnMap:          make(map[string]*daxConnEntry),
		//daxConnMap:          om.New[string, DaxConn](),
	}
	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		base.daxSrcEntryMap[ent.name] = ent
	}
	return &base
}

func (base *daxBaseImpl) SetUpLocalDaxSrc(name string, ds DaxSrc) Err {
	if base.isLocalDaxSrcsFixed {
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

	ent := &daxSrcEntry{name: name, ds: ds, local: true}
	if base.localDaxSrcList.head == nil {
		base.localDaxSrcList.head = ent
		base.localDaxSrcList.last = ent
	} else {
		ent.prev = base.localDaxSrcList.last
		base.localDaxSrcList.last.next = ent
		base.localDaxSrcList.last = ent
	}

	base.daxSrcEntryMap[ent.name] = ent

	return Ok()
}

func (base *daxBaseImpl) FreeLocalDaxSrc(name string) {
	if base.isLocalDaxSrcsFixed {
		return
	}

	ent := base.daxSrcEntryMap[name]
	if ent != nil && ent.local {
		delete(base.daxSrcEntryMap, name)
		if ent.prev != nil {
			ent.prev.next = ent.next
		}
		if ent.next != nil {
			ent.next.prev = ent.prev
		}
		ent.ds.End()
	}
}

func (base *daxBaseImpl) FreeAllLocalDaxSrcs() {
	if base.isLocalDaxSrcsFixed {
		return
	}

	for ent := base.localDaxSrcList.head; ent != nil; ent = ent.next {
		ent.ds.End()
	}
	base.localDaxSrcList.head = nil
	base.localDaxSrcList.last = nil
}

func (base *daxBaseImpl) begin() {
	base.isLocalDaxSrcsFixed = true
	isGlobalDaxSrcsFixed = true
}

func (base *daxBaseImpl) commit() Err {
	var wg sync.WaitGroup

	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		err := ent.conn.Commit(&wg)
		if err.IsNotOk() {
			wg.Wait()
			return NewErr(FailToCommitDaxConn{})
		}
	}

	wg.Wait()

	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		err := ent.conn.Committed()
		if err.IsNotOk() {
			return NewErr(FailToCommitDaxConn{})
		}
	}

	return Ok()
}

func (base *daxBaseImpl) rollback() {
	var wg sync.WaitGroup

	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		ent.conn.Rollback(&wg)
	}

	wg.Wait()
}

func (base *daxBaseImpl) end() {
	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		ent.conn.Close()
	}

	base.daxConnList.head = nil
	base.daxConnList.last = nil
	base.daxConnMap = make(map[string](*daxConnEntry))

	base.isLocalDaxSrcsFixed = false
}

func (base *daxBaseImpl) getDaxConn(name string) (DaxConn, Err) {
	ent, loaded := base.daxConnMap[name]
	if loaded {
		return ent.conn, Ok()
	}

	base.daxConnMutex.Lock()
	defer base.daxConnMutex.Unlock()

	dsEnt, exists := base.daxSrcEntryMap[name]
	if !exists {
		return nil, NewErr(DaxSrcIsNotFound{Name: name})
	}
	conn, err := dsEnt.ds.CreateDaxConn()
	if err.IsNotOk() {
		return nil, NewErr(FailToCreateDaxConn{Name: name}, err)
	}
	if conn == nil {
		return nil, NewErr(CreatedDaxConnIsNil{Name: name})
	}

	ent = &daxConnEntry{name: name, conn: conn}
	if base.daxConnList.head == nil {
		base.daxConnList.head = ent
		base.daxConnList.last = ent
	} else {
		ent.prev = base.daxConnList.last
		base.daxConnList.last.next = ent
		base.daxConnList.last = ent
	}

	base.daxConnMap[ent.name] = ent

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
