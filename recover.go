package main

import "fmt"

// this function panics
func mayPanic() {
	panic("a problem")
}

func main() {
	// recover must be called within a deferred function. when the
	// enclosing function panics, the defer will activate and a recover
	// call within it will catch the panic
	defer func() {
		// the return value of recover is the error raised in the call to
		// panic
		if r := recover(); r != nil {
			fmt.Println("recovered. Error:\n", r)
		}
	}()

	mayPanic()
	// this code will not run because mayPanic panics. the execution of
	// main stops at the point of the panic and resumes in the deferred
	// closure
	fmt.Println("after mayPanic()")
}
