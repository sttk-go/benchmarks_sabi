package ag_setup_noerr_sync_test

import (
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/ag_setup_noerr_sync"
)

type FooDaxConn struct{ err sabi.Err }

//type FooDaxConn struct{}

func (conn *FooDaxConn) Commit(ag *sabi.AsyncGroup) {}
func (conn *FooDaxConn) Committed() sabi.Err        { return conn.err }

// func (conn *FooDaxConn) Committed() sabi.Err          { return sabi.Ok() }
func (conn *FooDaxConn) Rollback(ag *sabi.AsyncGroup) {}
func (conn *FooDaxConn) Close()                       {}

type FooDaxSrc struct{ err sabi.Err }

//type FooDaxSrc struct{}

func (ds *FooDaxSrc) SetUp(ag *sabi.AsyncGroup) {}
func (ds *FooDaxSrc) Ready() sabi.Err           { return ds.err }

// func (ds *FooDaxSrc) Ready() sabi.Err           { return sabi.Ok() }
func (ds *FooDaxSrc) End() {}
func (ds *FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return &FooDaxConn{}, sabi.Ok()
}
