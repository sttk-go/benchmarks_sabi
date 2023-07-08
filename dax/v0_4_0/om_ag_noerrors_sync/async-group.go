package om_ag_noerrors_sync

import (
	"sync"
)

type AsyncGroup interface {
	Add(fn func() Err)
	wait()
}

type asyncGroupAsync struct {
	wg    sync.WaitGroup
	errs  map[string]Err
	mutex sync.Mutex
	name  string
}

func (ag *asyncGroupAsync) Add(fn func() Err) {
	ag.wg.Add(1)
	go func(name string) {
		defer ag.wg.Done()
		err := fn()
		if err.IsNotOk() {
			ag.mutex.Lock()
			defer ag.mutex.Unlock()
			ag.errs[name] = err
		}
	}(ag.name)
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
