package asyncgroup_sync

import (
	om "github.com/sttk/orderedmap"
)

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcMap = om.New[string, DaxSrc]()
}

func AddGlobalDaxSrcForTest(name string, ds DaxSrc) {
	globalDaxSrcMap.Store(name, ds)
}

func FreeGlobalDaxSrcForTest(name string) {
	globalDaxSrcMap.Delete(name)
}

func AddLocalDaxSrcForTest(base DaxBase, name string, ds DaxSrc) {
	base.(*daxBaseImpl).localDaxSrcMap.Store(name, ds)
}

func FreeLocalDaxSrcForTest(base DaxBase, name string) {
	base.(*daxBaseImpl).localDaxSrcMap.Delete(name)
}

func FreeAllLocalDaxSrcsForTest(base DaxBase) {
	base.(*daxBaseImpl).localDaxSrcMap = om.New[string, DaxSrc]()
}
