package dax3_ready_async

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcMap = make(map[string]DaxSrc)
}

func AddLocalDaxSrcForTest(base DaxBase, name string, ds DaxSrc) {
	base.(*daxBaseImpl).localDaxSrcMap[name] = ds
}
