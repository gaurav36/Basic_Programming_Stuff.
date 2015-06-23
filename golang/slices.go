/* Slices : Similar to array but much more powerful then array
 * [1]. To create an empty slice with non-zero length, use the builtin make.
 *      we make a slice of strings of length 3 (initially zero-valued).
 * [2]. We can set and get just like with arrays.
 * [3]. "append", which returns a slice containing one or more new values.
 * [4]. Slices can also be copyed
 * [5]. Slice support a "slice" operator with the syntax slice[low:high].
 *      for example s[low:high] means range of element including index "low"
 *      to excluding "high".
 * [6]. Slices can be composed into multi-dimensional data structures.
 *      The length of the inner slices can vary, unlike with multi-dimensional
 *      arrays.
 *
 *      for more info about slice:
 *                        http://blog.golang.org/go-slices-usage-and-internals
 */

package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println ("slice s is: ", s)
	fmt.Println ("length of slice s is: ", len(s))

	/* [2] */
	 s[0] = "g"
	 s[1] = "o"
	 s[2] = "l"
	 fmt.Println ("slice s is:", s)

	 /* [3] */
	 s = append (s, "a")
	 s = append (s, "n", "g")
	 fmt.Println ("Slice s is:", s)

	 /* [4] */
	 c := make([]string, len(s))
	 copy (c,s)
	 fmt.Println ("copy of slice s, c is: ", c)

	 /* [5] */
	 l := s[2:5]
	 fmt.Println ("slicing of s, l is: ", l)

	 /* range from starting to 4 (don't include 5)*/
	 l = s[:5]
	 fmt.Println("sl2:", l)

	 /* range from 2nd index to last */
	 l = s[2:]
	 fmt.Println("sl3:", l)

	 t := []string{"g", "h", "i"}
	 fmt.Println("dcl:", t)

	 /* [6] */
	 twoD := make([][]int, 3)
	 for i:=0; i<3; i++ {
	 	innerlen := i +1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println ("towD slice is: ", twoD)
}
