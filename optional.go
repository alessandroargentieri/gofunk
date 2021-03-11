package main

import (
	"reflect"
)

// Optional struct implements the functional logic in Golang: il allows to wrap a value and define a series of
// transformations without worrying about the errors.
// At the end of the chain you can get the output value or the error, indipendently of which has been the step
// it has occurred.
type Optional struct {
	value interface{}
	err   error
}

// OptionalOf static function creates an Optional struct of some value, wrapping the one provided in input
func OptionalOf(value interface{}) Optional {
	return Optional{value, nil}
}

// Map function represent the single transformation step for the Optional.value
func (opt Optional) Map(f Function) Optional {
	if opt.err != nil {
		return opt
	}
	v, err := f(opt.value)
	return Optional{v, err}
}

// IsError function returns if an error has occurred anywhere in the transformation chain
func (opt Optional) IsError() bool {
	return opt.err != nil
}

// IsNil function returns true if Optional.value is a nil interface{}
func (opt Optional) IsNil() bool {
	return opt.value == nil

}

// Get function returns the Optional.value wrapped
func (opt Optional) Get() interface{} {
	return opt.value
}

// Error function returns the Optional.err wrapped
func (opt Optional) Error() interface{} {
	return opt.err
}

// GetOrElse function returns the Optional.value if no error occurred, else the defaultValue
func (opt Optional) GetOrElse(defaultValue interface{}) interface{} {
	if opt.err == nil {
		return opt.value
	}
	return defaultValue
}

// Type function reveals the type of the Optional.value wrapped
func (opt Optional) Type() interface{} {
	return reflect.TypeOf(opt.value)
}

