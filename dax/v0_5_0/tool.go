package v0_5_0

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcList.head = nil
	globalDaxSrcList.last = nil
}

func AddGlobalDaxSrcForTest(name string, ds DaxSrc) {
	ent := &daxSrcEntry{name: name, ds: ds}
	if globalDaxSrcList.head == nil {
		globalDaxSrcList.head = ent
		globalDaxSrcList.last = ent
	} else {
		ent.prev = globalDaxSrcList.last
		globalDaxSrcList.last.next = ent
		globalDaxSrcList.last = ent
	}
}

func FreeGlobalDaxSrcForTest(name string) {
	for ent := globalDaxSrcList.head; ent != nil; ent = ent.next {
		if ent.name == name {
			ent.prev.next = ent.next
			ent.next.prev = ent.prev
			ent.ds.Close()
			return
		}
	}
}

func AddLocalDaxSrcForTest(daxBase DaxBase, name string, ds DaxSrc) {
	ent := &daxSrcEntry{name: name, ds: ds}
	base := daxBase.(*daxBaseImpl)
	if base.localDaxSrcList.head == nil {
		base.localDaxSrcList.head = ent
		base.localDaxSrcList.last = ent
	} else {
		ent.prev = base.localDaxSrcList.last
		base.localDaxSrcList.last.next = ent
		base.localDaxSrcList.last = ent
	}
}

func FreeLocalDaxSrcForTest(daxBase DaxBase, name string) {
	base := daxBase.(*daxBaseImpl)
	for ent := base.localDaxSrcList.head; ent != nil; ent = ent.next {
		if ent.name == name {
			ent.prev.next = ent.next
			ent.next.prev = ent.prev
			ent.ds.Close()
			return
		}
	}
}

func FreeAllLocalDaxSrcsForTest(daxBase DaxBase) {
	base := daxBase.(*daxBaseImpl)
	base.localDaxSrcList.head = nil
	base.localDaxSrcList.last = nil
}

func Begin(base DaxBase) {
	base.begin()
}
