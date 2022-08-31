package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	next := in
	for _, stage := range stages {
		if stage == nil {
			continue
		}
		next = terminator(done, stage(next))
	}
	return next
}

func terminator(terminate, in In) Out {
	next := make(Bi)

	go func() {
		defer close(next)
		for {
			select {
			case value, ok := <-in:
				if !ok {
					return
				}
				next <- value
			case <-terminate:
				return
			}
		}
	}()
	return next
}
