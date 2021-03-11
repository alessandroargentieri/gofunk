package gofunk

import (
	"fmt"
	"testing"
)

func TestOptional(t *testing.T) {

	opt := OptionalOf(0).
		Map(addOne).              //1
		Map(addOne).              //2
		Map(addOne).              //3
		Map(addOneAndThrowError). //4
		Map(addOne).              //5
		Map(addOne)               //6

	isErr := opt.IsError()          // true
	err := opt.Error()              // Oh man!
	value := opt.Get().(int)        // 4
	final := opt.GetOrElse(0).(int) //0

	if isErr == false {
		t.Errorf("Expected error!")
	}

	if err == nil {
		t.Errorf("Expected error!")
	}

	if value != 4 {
		t.Errorf("Expected 4!")
	}

	if final != 0 {
		t.Errorf("Expected 0!")
	}

}

func addOne(input interface{}) (interface{}, error) {
	value := input.(int)
	value = value + 1
	return value, nil
}

func addOneAndThrowError(input interface{}) (interface{}, error) {
	value := input.(int)
	value = value + 1
	return value, fmt.Errorf("Oh man!")
}
