package ready_sync

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcMap = make(map[string]DaxSrc)
}

func AddGlobalDaxSrcForTest(name string, ds DaxSrc) {
	globalDaxSrcMap[name] = ds
}

func AddLocalDaxSrcForTest(base DaxBase, name string, ds DaxSrc) {
	base.(*daxBaseImpl).localDaxSrcMap[name] = ds
}
