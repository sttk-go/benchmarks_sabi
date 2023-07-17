package v0_5_0

import (
	"sync"
)

type AsyncGroup interface {
	Add(fn func() Err)
	wait()
}

type errEntry struct {
	name string
	err  Err
	next *errEntry
}

type asyncGroupAsync struct {
	wg sync.WaitGroup
	//errs  map[string]Err
	errHead *errEntry
	errLast *errEntry
	mutex   sync.Mutex
	name    string
}

func (ag *asyncGroupAsync) Add(fn func() Err) {
	ag.wg.Add(1)
	go func(name string) {
		defer ag.wg.Done()
		err := fn()
		if err.IsNotOk() {
			ag.mutex.Lock()
			defer ag.mutex.Unlock()
			//ag.errs[name] = err
			ag.addErr(name, err)
		}
	}(ag.name)
}

func (ag *asyncGroupAsync) addErr(name string, err Err) {
	ent := &errEntry{name: name, err: err}
	if ag.errLast == nil {
		ag.errHead = ent
		ag.errLast = ag.errHead
	} else {
		ag.errLast.next = ent
		ag.errLast = ag.errLast.next
	}
}

func (ag *asyncGroupAsync) hasErr() bool {
	return (ag.errHead != nil)
}

func (ag *asyncGroupAsync) makeErrs() map[string]Err {
	m := make(map[string]Err)
	for ent := ag.errHead; ent != nil; ent = ent.next {
		m[ent.name] = ent.err
	}
	return m
}

func (ag *asyncGroupAsync) wait() {
	ag.wg.Wait()
}

type asyncGroupSync struct {
	err Err
}

func (ag *asyncGroupSync) Add(fn func() Err) {
	ag.err = fn()
}

func (ag *asyncGroupSync) wait() {}
