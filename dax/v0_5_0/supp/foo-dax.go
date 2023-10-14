package supp_test

import (
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit(ag sabi.AsyncGroup) sabi.Err { return sabi.Ok() }
func (conn FooDaxConn) Rollback(ag sabi.AsyncGroup)        {}
func (conn FooDaxConn) Close()                             {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) Setup(ag sabi.AsyncGroup) sabi.Err { return sabi.Ok() }
func (ds FooDaxSrc) Close()                            {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}
