package ready_sync

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcMap = make(map[string]DaxSrc)
}

func AddGlobalDaxSrcForTest(name string, ds DaxSrc) {
	globalDaxSrcMap[name] = ds
}

func FreeGlobalDaxSrcForTest(name string) {
	delete(globalDaxSrcMap, name)
}

func AddLocalDaxSrcForTest(base DaxBase, name string, ds DaxSrc) {
	base.(*daxBaseImpl).localDaxSrcMap[name] = ds
}

func FreeLocalDaxSrcForTest(base DaxBase, name string) {
	delete(base.(*daxBaseImpl).localDaxSrcMap, name)
}

func FreeAllLocalDaxSrcsForTest(base DaxBase) {
	base.(*daxBaseImpl).localDaxSrcMap = make(map[string]DaxSrc)
}
