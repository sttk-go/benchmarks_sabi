package v0_6_0

func ResetGlobals() {
	isGlobalDaxSrcsFixed = false
	globalDaxSrcEntryList.head = nil
	globalDaxSrcEntryList.last = nil
}

func AddGlobalDaxSrcForTest(name string, ds DaxSrc) {
	ent := &daxSrcEntry{name: name, ds: ds}
	if globalDaxSrcEntryList.head == nil {
		globalDaxSrcEntryList.head = ent
		globalDaxSrcEntryList.last = ent
	} else {
		ent.prev = globalDaxSrcEntryList.last
		globalDaxSrcEntryList.last.next = ent
		globalDaxSrcEntryList.last = ent
	}
}

func FreeGlobalDaxSrcForTest(name string) {
	for ent := globalDaxSrcEntryList.head; ent != nil; ent = ent.next {
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
	if base.localDaxSrcEntryList.head == nil {
		base.localDaxSrcEntryList.head = ent
		base.localDaxSrcEntryList.last = ent
	} else {
		ent.prev = base.localDaxSrcEntryList.last
		base.localDaxSrcEntryList.last.next = ent
		base.localDaxSrcEntryList.last = ent
	}
}

func FreeLocalDaxSrcForTest(daxBase DaxBase, name string) {
	base := daxBase.(*daxBaseImpl)
	for ent := base.localDaxSrcEntryList.head; ent != nil; ent = ent.next {
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
	base.localDaxSrcEntryList.head = nil
	base.localDaxSrcEntryList.last = nil
}

func Begin(base DaxBase) {
	base.begin()
}
