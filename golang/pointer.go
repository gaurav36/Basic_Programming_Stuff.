/* Pointer in go: pointer in go is same as C/C++ porgramming language
 */

package main

import (
	"fmt"
)

func zeroval (ival int) {
	ival = 0
}

func zeroptr (iptr *int) {
	*iptr = 0
}

func main() {
	i := 2
	fmt.Println ("Initially i is: ", i)

	zeroval (i)
	fmt.Println ("zeroval i is: ", i)

	zeroptr (&i)
	fmt.Println ("zeroptr i is: ", i)
}
