package asyncgroup_sync_test

import (
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/asyncgroup_sync"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit(ag *sabi.AsyncGroup) sabi.Err {
	return sabi.Ok()
}
func (conn FooDaxConn) Rollback(ag *sabi.AsyncGroup) {}
func (conn FooDaxConn) Close()                       {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) SetUp(ag *sabi.AsyncGroup) sabi.Err { return sabi.Ok() }
func (ds FooDaxSrc) End()                               {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}
