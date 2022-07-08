package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	taskChannel := make(chan Task)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	errorCounter := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskChannel {
				err := task()
				if err != nil {
					mu.Lock()
					errorCounter++
					mu.Unlock()
				}
			}
		}()
	}

	for _, task := range tasks {
		if errorCounter >= m {
			break
		}
		taskChannel <- task
	}
	close(taskChannel)
	wg.Wait()

	if errorCounter >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
