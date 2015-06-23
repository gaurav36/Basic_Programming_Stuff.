/* Goroutines: Exciting feature of go language. Go routines is a lightweight
 * thread of execution.
 * Goroutines being run concurrently by the Go runtime.
 * [1]. normally function will execute synchronously. goroutines accept
 *      to execute function asynchronously in seperate goroutines, means
 *      two function are executing parallelly.
 * [2]. You can also start a goroutine for an anonymous function call.
 */

package main

import (
	"fmt"
)

func f(arg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(arg, ":", i)
	}
}

func main() {

	f("direct")

	/* [1] */
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("hello golang")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
