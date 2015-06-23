/* Panic in Go: A panic typically means something went unexpectedly wrong.
 *              Mostly we use it to fail fast on errors that shouldn’t occur
 *              during normal operation, or that we aren’t prepared to handle
 *              gracefully.
 * Use case of panic: A common use of panic is to abort if a function returns
 *                    an error value that we don’t know how to (or want to)
 *                    handle. 
 */

package main

import "os"

func main() {
        panic ("a problem")

	_, err := os.Create ("/tmp/panic_go")
	if err != nil {
		panic (err)
	}
}
