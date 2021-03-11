package gofunk

// Future struct is a monad implementing a parallel task to be performed
type Future struct {
	ch          *chan Either
	fn          Function
	input       interface{}
	output      *Either
	isCompleted *bool
}

// FutureFrom static function generates a Future object from a function to be implemented asyncronously and its input
func FutureFrom(fn Function, input interface{}) Future {
	ch := make(chan Either, 1)
	isComplete := false
	return Future{
		ch:          &ch,
		fn:          fn,
		input:       input,
		output:      nil,
		isCompleted: &isComplete,
	}
}

// ProcessAsync function performs the wrapped function in another Goroutine and returns a Future with the wrapped result
// It can also used as a "void" because the wrapped chan is pointed by a Pointer
func (future Future) ProcessAsync() Future {
	if future.isCompleted != nil && *future.isCompleted == true {
		return future
	}
	go channelifyProcess(future.fn, future.input, future.ch)
	isComplete := true
	*future.isCompleted = isComplete
	return Future{
		ch:          future.ch,
		fn:          future.fn,
		input:       future.input,
		output:      future.output,
		isCompleted: &isComplete,
	}

}

func channelifyProcess(fn Function, input interface{}, ch *chan Either) {
	output, err := fn(input)
	if err != nil {
		*ch <- EitherFromError(err)
		return
	}
	*ch <- EitherFromResult(output)
	return
}

// BlockingReturn waits and gets the result to the main Goroutine in the form of an Either object
func (future *Future) BlockingReturn() Either {
	if future.output != nil {
		return *future.output
	}
	either := <-*future.ch
	future.output = &either
	//*future.output = either // va in panic
	return either
}
