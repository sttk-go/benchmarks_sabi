package dax3_ready_async

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcMap = make(map[string]DaxSrc)
}
