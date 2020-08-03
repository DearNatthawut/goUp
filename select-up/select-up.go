package selectup

import (
	"fmt"
	"time"
)

// DoLongWork : do somting 3 second
func DoLongWork() int {
	time.Sleep(3 * time.Second)
	return 1
}

// DoWorkWithLimitTime : do somting with limit time
func DoWorkWithLimitTime(d time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		ch <- DoLongWork()
	}()

	select {
	case r := <-ch:
		return r, nil
	case <-time.After(d):
		return 0, fmt.Errorf("timeout")
	}
}
