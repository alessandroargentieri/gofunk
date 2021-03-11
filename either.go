package gofunk

// Either struct is the equivalent of functional Scala Either: a wrapper for the result/error response
type Either struct {
	result interface{}
	err    error
}

// EitherFromResult static function initializes a Either object with a result
func EitherFromResult(result interface{}) Either {
	return Either{result, nil}
}

// EitherFromError static function initializes a Either object with a
func EitherFromError(err error) Either {
	return Either{nil, err}
}

// IsError function returns true if Either wraps an error object
func (e Either) IsError() bool {
	return e.err != nil
}

// IsResult function returns true if Either doesn't wrap an error object
func (e Either) IsResult() bool {
	return e.err == nil
}

// GetResult function returns the wrapped Result
func (e Either) GetResult() interface{} {
	if e.err != nil {
		panic("Either struct contains an error")
	}
	return e.result
}

// GetError function returns the wrapped error if any
func (e Either) GetError() error {
	return e.err
}

// Get function return the wrapped result which can "either" be an error or another object
func (e Either) Get() interface{} {
	if e.IsError() {
		return e.err
	}
	return e.result
}

// GetOrElse function returns the wrapped result or if an error is present, a default value
func (e Either) GetOrElse(defaultResult interface{}) interface{} {
	if e.IsError() {
		return defaultResult
	}
	return e.result
}
