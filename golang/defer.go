/* Defer in Go: Defer is used to ensure that a function call is performed later
 *              in a program’s execution, usually for purposes of cleanup.
 *
 * Use case of Defer: Suppose we wanted to create a file, write to it,
 * and then close when we’re done. so we can use defer just after creating
 * of file.
 */

package main

import (
	"fmt"
	"os"
)

func test2 () {
	fmt.Println ("three")
}

func test () {
	fmt.Println ("one")
	fmt.Println ("two")
	defer test2()
	fmt.Println ("four")
	fmt.Println ("five")
}

func main() {
	f := createFile ("/tmp/defer.txt")
	defer closeFile (f) /* this function execute at end of main*/
	writeFile (f)
	
	test ();
}

func createFile(p string) *os.File {
	fmt.Println ("Creating file...")

	f, err := os.Create(p)
	if err != nil {
		panic (err)
	}
	return f
}

func writeFile (f *os.File) {
	fmt.Println ("Writing in to file...")
	fmt.Fprintln (f, "defer in Go")
	fmt.Fprintln (f, "my data")
}

func closeFile (f *os.File) {
	fmt.Println ("Closing file...")
	f.Close ()
}
