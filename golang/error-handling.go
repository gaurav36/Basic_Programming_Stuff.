/* Errors handling in Go:
 * [1]. By convention, errors are the last return value and have type error,
 *      a built-in interface.
 * [2]. A nil value in the error position indicates that there was no error.
 * [3]. errors.New constructs a basic error value with the given error message.
 * [4]. It’s possible to use custom types as errors by implementing the Error()
 *      method on them.
 * for more details on error handling: http://blog.golang.org/error-handling-and-go
 */

package main

import (
	"errors" /*built in error package*/
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant't work with 42")
	}

	/* [2] */
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}
}
