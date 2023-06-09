package dax3_ready_async_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax3_ready_async"
	"sync"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit() sabi.Err { return sabi.Ok() }
func (conn FooDaxConn) Rollback()        {}
func (conn FooDaxConn) Close()           {}

type FooDaxSrc struct{ err sabi.Err }

func (ds FooDaxSrc) SetUp(wg sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return ds.err
}
func (ds FooDaxSrc) Ready() sabi.Err {
	return ds.err
}
func (ds FooDaxSrc) End(wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}

type BarDaxConn struct{}

func (conn BarDaxConn) Commit() sabi.Err { return sabi.Ok() }
func (conn BarDaxConn) Rollback()        {}
func (conn BarDaxConn) Close()           {}

type BarDaxSrc struct{ err sabi.Err }

func (ds BarDaxSrc) SetUp(wg sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return ds.err
}
func (ds BarDaxSrc) Ready() sabi.Err {
	return ds.err
}
func (ds BarDaxSrc) End(wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
}
func (ds BarDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return BarDaxConn{}, sabi.Ok()
}

type BazDaxConn struct{}

func (conn BazDaxConn) Commit() sabi.Err { return sabi.Ok() }
func (conn BazDaxConn) Rollback()        {}
func (conn BazDaxConn) Close()           {}

type BazDaxSrc struct{ err sabi.Err }

func (ds BazDaxSrc) SetUp(wg sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return ds.err
}
func (ds BazDaxSrc) Ready() sabi.Err {
	return ds.err
}
func (ds BazDaxSrc) End(wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
}
func (ds BazDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return BazDaxConn{}, sabi.Ok()
}

type QuxDaxConn struct{}

func (conn QuxDaxConn) Commit() sabi.Err { return sabi.Ok() }
func (conn QuxDaxConn) Rollback()        {}
func (conn QuxDaxConn) Close()           {}

type QuxDaxSrc struct{ err sabi.Err }

func (ds QuxDaxSrc) SetUp(wg sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return ds.err
}
func (ds QuxDaxSrc) Ready() sabi.Err {
	return ds.err
}
func (ds QuxDaxSrc) End(wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
}
func (ds QuxDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return QuxDaxConn{}, sabi.Ok()
}

type CorgeDaxConn struct{}

func (conn CorgeDaxConn) Commit() sabi.Err { return sabi.Ok() }
func (conn CorgeDaxConn) Rollback()        {}
func (conn CorgeDaxConn) Close()           {}

type CorgeDaxSrc struct{ err sabi.Err }

func (ds CorgeDaxSrc) SetUp(wg sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return ds.err
}
func (ds CorgeDaxSrc) Ready() sabi.Err {
	return ds.err
}
func (ds CorgeDaxSrc) End(wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
}
func (ds CorgeDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return CorgeDaxConn{}, sabi.Ok()
}
