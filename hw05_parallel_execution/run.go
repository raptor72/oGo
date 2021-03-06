package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	taskChan := make(chan Task)
	m32 := int32(m)
	wg := new(sync.WaitGroup)
	var errorsCount int32
	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for task := range taskChan {
				if atomic.LoadInt32(&errorsCount) >= m32 {
					return
				}
				err := task()
				if err != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		if atomic.LoadInt32(&errorsCount) >= m32 {
			break
		}
		taskChan <- task
	}
	close(taskChan)
	wg.Wait()
	if errorsCount >= m32 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
