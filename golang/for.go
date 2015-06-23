/* for is Go's only looping construct.
 * for without a condition will loop repeatedly until you break out of the loop
 */

package main

import (
	"fmt"
//	"math"   /* if i am importing math package and not using in programme then go will through an error*/
)

func main() {
	i :=1
	
	for i <= 3 {
		fmt.Println(i)
		i = i +1
	}

	for j := 7; j <=9; j++ {
		fmt.Println (j)
        }

	for {
		fmt.Println ("loop")
		break
        }
}

