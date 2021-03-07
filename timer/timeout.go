package timer

import (
	"fmt"
	"time"
)

var (
	errCallFnTimeout = fmt.Errorf("call fn timeout")
)

type timeoutFn func() error

// CallFnWithTimeout ...
func CallFnWithTimeout(fn timeoutFn, d time.Duration) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- fn()
	}()

	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		return errCallFnTimeout
	case err := <-errChan:
		fmt.Println("call fn err: ", err)
		return err
	}
}
