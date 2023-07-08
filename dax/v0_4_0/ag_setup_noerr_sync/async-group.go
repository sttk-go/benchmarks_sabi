package ag_setup_noerr_sync

import (
	"sync"
)

type AsyncGroup struct {
	wg    sync.WaitGroup
	not   bool
	isErr bool
}

func (ag *AsyncGroup) Add(fn func() Err) {
	if ag.not {
		fn()
		return
	}
	ag.wg.Add(1)
	go func() {
		defer ag.wg.Done()
		fn()
	}()
}

func (ag *AsyncGroup) wait() {
	if !ag.not {
		ag.wg.Wait()
	}
}
