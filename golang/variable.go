/* Variables in go:
 * [1]. var declare 1 or more variable. you can declare or initialize
 *      variable after importing "fmt" package.
 * [2]. Variables declared without a corresponding initialization are 
 *      zero-valued.
 * [3]. ":=" syntax is for declaring and initializing a variable.
 * [4]. Go will automatically take care of data type of the variable.
 *      you do not need to give data type of variable while declaration
 *      as you give's in other programming language.
 */

package main

//var temp int /* wrong!! you can't declare variable before importing*/

import (
        "fmt"
)

func main() {

	/* There is no character in go language*/
	/* It will treat ch as a integer. and assign ASCII value of A (65) to ch variable */
	var ch = 'A'
	fmt.Println ("ch is: ", ch)

	var a string = "initial"
	fmt.Println (a)

	var b, c int = 1,3
	fmt.Println("b is", b, "c is", c)

	var d int
        fmt.Println ("d is: ", d)

	e := "string"
	fmt.Println ("e is: ", e);

	const x string = "hello constant"
	fmt.Println ("const x is: ", x);

	/* %v formate specifier support all data type (float, char, int, string)*/
	fmt.Printf ("constant x is: %v\n",x)

	/* you can't change the value of constant variable*/
	//x = "trying to change constant"
	//fmt.Println ("updated constant x is:", x)
}
