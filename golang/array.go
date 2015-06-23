/* Array example in go language:
 * [1]. By default array will initialize with zero if you are using var keyword
 * [2]. We can set a value at an index using the array[index] = value syntax.
 * [3]. The builtin len returns the length of an array.
 */

 package main

 import "fmt"

 func main() {
        /* [1] */
 	var a [5]int
	fmt.Println("Array a is: ", a)

        /* [2] */
	a[4] = 100
	fmt.Println ("Array a is", a)

	/* [3] */
	fmt.Println ("Length of array a is: ", len(a))

	b := [5]int{} /* we can't simply use b := [5]int  we need to initialize
	               * array with something*/
	fmt.Println ("Array b is: ", b)


	c := [5]int{1,2,3,4,5}
	fmt.Println ("Array c is: ", c)


	var twoD [2][3]int
	for i :=0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println ("twoD array is: ", twoD);

	d := [4]string{"a", "b", "c", "d"}
	fmt.Println ("Array d is: ", d)


}
