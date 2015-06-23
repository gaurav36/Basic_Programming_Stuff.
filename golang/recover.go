/* panic and recover: like try and catch in C++ programming.
 */

package main

import "fmt"


func temp1() {
	fmt.Println ("In function temp1")
}

func temp() {
	str := recover()
	fmt.Println (str)
	fmt.Println ("hello")
}

func main() {
	defer temp()
	temp1 ()
	panic ("PANIC")

	/* call to panic immediately stops execution of the function. it
	 * will not recover. instead we have to pair panic with defer.
	 * defer will gurantee that before stopping execution it will
	 * call appropriate defer function. and in that function we can
	 * recover panic.*/
       //	str := recover()
	//fmt.Println (str)

	fmt.Println ("recover")
	fmt.Println ("recover")
	fmt.Println ("recover")
	fmt.Println ("recover")
	
	temp1()
}
