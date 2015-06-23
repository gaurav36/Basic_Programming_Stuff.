/* range: range iterates over of elements in a variety of data structures. 
 * [1]. range on arrays and slices provides both the index and value for each
 *      entry.
 * [2]. range on map iterates over key/value pairs.
 * [3]. range on strings iterates over Unicode (ASCII value) code points.
 *      The first value is the starting byte index of the rune and the second
 *      the rune itself.
 */

package main

import (
 	"fmt"
)

func main() {
	nums := []int {1, 2, 3, 4, 5}
	fmt.Println ("array nums is: ", nums)

	sum := 0

	/* [1] */
	/* here we don't need index value so we are ignoring index value
	 * by "_"  . If we need index value then we can have i, num := range nums
	 */
	for _, num := range nums {
		sum += num
	}
	fmt.Println ("sum of element's of array nums is: ", sum)

	/* Printing index value of array*/
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index is:", i)
		}
	}

	/* [2] */
	kvs := map[string]string{"a": "apple", "b": "boy"}
	for k, v := range kvs {
		fmt.Printf("%s --> %s\n",k,v)
	}

	/* [3] */
	for i, c := range "go" {  /* ASCII value of g is 103 and o is 111*/
		fmt.Println ("i is: ", i, "c is: ", c)
	}
	
}

