/* Command line argument: go run hello.go uses run and hello.go arguments to the go program.
 * [1]. os.Args provides access to raw command-line arguments. Note that the
 *      first value in this slice is the path to the program, and os.
 *      Args[1:] holds the arguments to the program.
 * [2]. You can get individual args with normal indexing.
 */ 

package main

import (
 	"fmt"
	"os"
)

func main() {
	args_with_prog := os.Args
	args_without_prog := os.Args [1:]


	args := os.Args[3]


	fmt.Println ("args with prog is",args_with_prog)
	fmt.Println ("arguments without binary execution is:",args_without_prog)

	fmt.Println ("args 3 is", args)
}

/* first build the programme using # go build file_name.go
 * then pass command line argument during execution of programme
 */
