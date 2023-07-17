package v0_5_0

import (
	"reflect"
	"sync"
)

type (
	FailToSetupGlobalDaxSrcs struct {
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

type DaxConn interface {
	Commit(ag AsyncGroup) Err
	Rollback(ag AsyncGroup)
	Close()
}

type DaxSrc interface {
	Setup(ag AsyncGroup) Err
	Close()
	CreateDaxConn() (DaxConn, Err)
}

type daxSrcEntry struct {
	name  string
	ds    DaxSrc
	local bool
	prev  *daxSrcEntry
	next  *daxSrcEntry
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
	isGlobalDaxSrcsFixed bool = false
	globalDaxSrcList     daxSrcList
)

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

func SetupGlobalDaxSrcs() Err {
	isGlobalDaxSrcsFixed = true

	var ag asyncGroupAsync

	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		ag.name = ent.name
		err := ent.ds.Setup(&ag)
		if err.IsNotOk() {
			ag.wait()
			CloseGlobalDaxSrcs()
			ag.addErr(ag.name, err)
			return NewErr(FailToSetupGlobalDaxSrcs{Errors: ag.makeErrs()})
		}
	}

	ag.wait()

	if ag.hasErr() {
		return NewErr(FailToSetupGlobalDaxSrcs{Errors: ag.makeErrs()})
	}

	return Ok()
}

func CloseGlobalDaxSrcs() {
	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		ent.ds.Close()
	}
}

type Dax interface {
	getDaxConn(name string) (DaxConn, Err)
}

type DaxBase interface {
	Dax
	AddLocalDaxSrc(name string, ds DaxSrc) Err
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

	daxConnList  daxConnList
	daxConnMap   map[string](*daxConnEntry)
	daxConnMutex sync.Mutex

	asyncGroup asyncGroupSync
}

func NewDaxBase() DaxBase {
	base := daxBaseImpl{
		isLocalDaxSrcsFixed: false,
		localDaxSrcList:     daxSrcList{},
		daxSrcEntryMap:      make(map[string]*daxSrcEntry),
		daxConnList:         daxConnList{},
		daxConnMap:          make(map[string]*daxConnEntry),
	}
	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		base.daxSrcEntryMap[ent.name] = ent
	}
	return &base
}

func (base *daxBaseImpl) AddLocalDaxSrc(name string, ds DaxSrc) Err {
	if base.isLocalDaxSrcsFixed {
		return Ok()
	}

	err := ds.Setup(&(base.asyncGroup))
	if err.IsNotOk() {
		return err
	}

	if base.asyncGroup.err.IsNotOk() {
		return base.asyncGroup.err
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
			ent.prev.prev = ent.prev
		}
		ent.ds.Close()
	}
}

func (base *daxBaseImpl) FreeAllLocalDaxSrcs() {
	if base.isLocalDaxSrcsFixed {
		return
	}

	for ent := base.localDaxSrcList.head; ent != nil; ent = ent.next {
		ent.ds.Close()
		delete(base.daxSrcEntryMap, ent.name)
	}
	base.localDaxSrcList.head = nil
	base.localDaxSrcList.last = nil
}

func (base *daxBaseImpl) begin() {
	base.isLocalDaxSrcsFixed = true
	isGlobalDaxSrcsFixed = true
}

func (base *daxBaseImpl) commit() Err {
	var ag asyncGroupAsync

	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		ag.name = ent.name
		err := ent.conn.Commit(&ag)
		if err.IsNotOk() {
			ag.wait()
			ag.addErr(ent.name, err)
			return NewErr(FailToCommitDaxConn{Errors: ag.makeErrs()})
		}
	}

	ag.wait()

	if ag.hasErr() {
		return NewErr(FailToCommitDaxConn{Errors: ag.makeErrs()})
	}

	return Ok()
}

func (base *daxBaseImpl) rollback() {
	var ag asyncGroupAsync

	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		ent.conn.Rollback(&ag)
	}

	ag.wait()
}

func (base *daxBaseImpl) end() {
	for ent := base.daxConnList.head; ent != nil; ent = ent.next {
		ent.conn.Close()
		delete(base.daxConnMap, ent.name)
	}

	base.daxConnList.head = nil
	base.daxConnList.last = nil

	base.isLocalDaxSrcsFixed = false
}

func (base *daxBaseImpl) getDaxConn(name string) (DaxConn, Err) {
	connEnt, loaded := base.daxConnMap[name]
	if loaded {
		return connEnt.conn, Ok()
	}

	base.daxConnMutex.Lock()
	defer base.daxConnMutex.Unlock()

	dsEnt, exists := base.daxSrcEntryMap[name]
	if !exists {
		return nil, NewErr(DaxSrcIsNotFound{Name: name})
	}
	conn, err := dsEnt.ds.CreateDaxConn()
	if err.IsNotOk() {
		return nil, NewErr(FailToCreateDaxConn{Name: name})
	}
	if conn == nil {
		return nil, NewErr(CreatedDaxConnIsNil{Name: name})
	}

	connEnt = &daxConnEntry{name: name, conn: conn}
	if base.daxConnList.head == nil {
		base.daxConnList.head = connEnt
		base.daxConnList.last = connEnt
	} else {
		connEnt.prev = base.daxConnList.last
		base.daxConnList.last.next = connEnt
		base.daxConnList.last = connEnt
	}

	base.daxConnMap[connEnt.name] = connEnt

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
		t = reflect.TypeOf(conn)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			from = "*" + t.Name() + " (" + t.PkgPath() + ")"
		} else {
			from = t.Name() + " (" + t.PkgPath() + ")"
		}

		err = NewErr(FailToCastDaxConn{Name: name, FromType: from, ToType: to})
	}

	return *new(C), err
}
