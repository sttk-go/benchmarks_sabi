package supp_test

import (
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"
	errs "github.com/sttk/benchmarks_sabi/dax/v0_6_0/errs"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit(ag sabi.AsyncGroup) errs.Err { return errs.Ok() }
func (conn FooDaxConn) IsCommitted() bool                  { return true }
func (conn FooDaxConn) Rollback(ag sabi.AsyncGroup)        {}
func (conn FooDaxConn) ForceBack(ag sabi.AsyncGroup)       {}
func (conn FooDaxConn) Close()                             {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) Setup(ag sabi.AsyncGroup) errs.Err { return errs.Ok() }
func (ds FooDaxSrc) Close()                            {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, errs.Err) {
	return FooDaxConn{}, errs.Ok()
}
