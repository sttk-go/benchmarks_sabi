package ready_async_test

import (
	"sync"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/ready_async"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit(wg *sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return sabi.Ok()
}
func (conn FooDaxConn) Committed() sabi.Err { return sabi.Ok() }
func (conn FooDaxConn) Rollback(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
}
func (conn FooDaxConn) Close() {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) SetUp(wg *sync.WaitGroup) sabi.Err {
	wg.Add(1)
	go func() {
		wg.Done()
	}()
	return sabi.Ok()
}
func (ds FooDaxSrc) Ready() sabi.Err { return sabi.Ok() }
func (ds FooDaxSrc) End()            {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}
