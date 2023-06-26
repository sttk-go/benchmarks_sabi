package orderedmap_isasync_async_test

import (
	"sync"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_async"
)

type FooDaxConn struct{}

func (conn FooDaxConn) IsAsync() bool { return true }
func (conn FooDaxConn) Commit(wg *sync.WaitGroup) sabi.Err {
	if wg != nil {
		wg.Add(1)
		go func() {
			wg.Done()
		}()
	}
	return sabi.Ok()
}
func (conn FooDaxConn) Committed() sabi.Err { return sabi.Ok() }
func (conn FooDaxConn) Rollback(wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
		go func() {
			wg.Done()
		}()
	}
}
func (conn FooDaxConn) Close() {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) IsAsync() bool { return true }
func (ds FooDaxSrc) SetUp(wg *sync.WaitGroup) sabi.Err {
	if wg != nil {
		wg.Add(1)
		go func() {
			wg.Done()
		}()
	}
	return sabi.Ok()
}
func (ds FooDaxSrc) Ready() sabi.Err { return sabi.Ok() }
func (ds FooDaxSrc) End()            {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}
