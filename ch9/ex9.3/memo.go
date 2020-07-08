package memo

import (
	"fmt"
)

// Func is the type of the function to memoize
type Func func(key string, done <- chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	done     <-chan struct{}
	response chan<- result
}

// Memo ...
type Memo struct{ requests, cancels chan request }

// New Memo
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), cancels: make(chan request)}
	go memo.server(f)
	return memo
}

// Get ...
func (memo *Memo) Get(key string, done <- chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, done,response}
	memo.requests <- req
	res := <-response

	select {
	case <- done:
		fmt.Println("get: queueing cancellation request")
		memo.cancels <- req
	default:
		// Not cancelled. Continue
	}
	return res.value, res.err
}

// Close ...
func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
Loop:

	for {
		Cancel:
			// handle all cancelled request
			for {
				select {
				case req := <-memo.cancels:
					fmt.Println("server: deleting cancelled entry (early)")
					delete(cache, req.key)
				default:
					break Cancel
				}
			}

		select {
		case req := <- memo.cancels:
			fmt.Println("server: deleting cancelled entry")
			delete(cache, req.key)
			continue Loop
		case req := <- memo.requests:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done <- chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
