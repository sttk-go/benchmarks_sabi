package asyncgroup_sync

import (
	"sync"
)

type AsyncGroup struct {
	wg    sync.WaitGroup
	not   bool
	errs  []Err
	isErr bool
	index int
}

func (ag *AsyncGroup) Add(fn func() Err) {
	if ag.not {
		err := fn()
		if err.IsNotOk() {
			ag.errs[ag.index] = err
			ag.isErr = true
		}
		return
	}
	ag.wg.Add(1)
	go func(i int) {
		defer ag.wg.Done()
		err := fn()
		if err.IsNotOk() {
			ag.errs[i] = err
			ag.isErr = true
		}
	}(ag.index)
}

func (ag *AsyncGroup) wait() {
	if !ag.not {
		ag.wg.Wait()
	}
}
