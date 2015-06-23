/* Multiple Return Values: go support multiple return value
 * [1]. The (int, int) in this function signature shows that the function returns 2
 * ints.
 * [2]. If you only want a subset of the returned values, use the blank
 * identifier "_"
 */

package main

import (
 	"fmt"
)

func vals () (int, int) {
	return 3, 7
}

func main() {
	/* [1] */
	a, b := vals()
	fmt.Println ("a is: ",a)
	fmt.Println ("b is: ",b)

	/* [2] */
	_, c := vals()
	fmt.Println ("c is: ",c)
}

