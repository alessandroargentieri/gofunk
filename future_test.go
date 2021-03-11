package gofunk

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	start := time.Now()

	f1 := FutureFrom(HeavyTask1, "first-task-input")
	f2 := FutureFrom(HeavyTask2, "second-task-input")
	f3 := FutureFrom(HeavyTask3, "third-task-input")

	f1.ProcessAsync()
	f2.ProcessAsync()
	f3.ProcessAsync()

	fmt.Println("Keep on the main Goroutine")
	fmt.Println(time.Since(start))
	if time.Since(start) > time.Second {
		t.Errorf("The processes must be started asyncronously")
	}

	res1 := f1.BlockingReturn().GetResult().(string)
	res2 := f2.BlockingReturn().GetResult().(string)
	res3 := f3.BlockingReturn().GetError()

	if res1 != "first-task: 1 DONE!" {
		t.Errorf("Wrong return for task 1")
	}

	if res2 != "second-task: 2 DONE!" {
		t.Errorf("Wrong return for task 2")
	}

	if res3.Error() != "Error task3" {
		t.Errorf("Wrong return for task 3")
	}

	if time.Since(start) < time.Duration(time.Duration.Seconds(2)) {
		t.Errorf("The processes must be performed in 5 seconds (max duration of the three tasks)")
	}
}

func TestFutureIdempotence(t *testing.T) {
	f1 := FutureFrom(HeavyTask1, "first-task-input")
	f1.ProcessAsync().ProcessAsync().ProcessAsync()
	f1.ProcessAsync()

	f1.BlockingReturn()
	f1.BlockingReturn()

	res1 := f1.BlockingReturn().GetResult().(string)

	if res1 != "first-task: 1 DONE!" {
		t.Errorf("Wrong return for task 1")
	}

}

func HeavyTask1(input interface{}) (interface{}, error) {
	time.Sleep(5 * time.Second)
	return strings.Split(input.(string), "-input")[0] + ": 1 DONE!", nil
}

func HeavyTask2(input interface{}) (interface{}, error) {
	time.Sleep(3 * time.Second)
	return strings.Split(input.(string), "-input")[0] + ": 2 DONE!", nil
}

func HeavyTask3(input interface{}) (interface{}, error) {
	time.Sleep(2 * time.Second)
	return nil, fmt.Errorf("Error task3")
}
