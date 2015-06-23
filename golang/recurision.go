/* Recurision in go
 */

package main

import (
 	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n*fact(n-1)
}

func main() {
	var input int
	fmt.Println ("Enter value")
	fmt.Scanln (&input)
	fmt.Printf ("factorial of %d is %d\n", input, fact(input))
}
