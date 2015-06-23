/* Variable argument function in golang
 */

package main

import (
	"fmt"
)

/* we can assume that here variable argument treated as a array*/
func core(varg ...int) {
	fmt.Println(varg)
	//fmt.Println(varg, "")
	fmt.Println("length of variable argument is: ", len(varg))

	total := 0
	for _, num := range varg {
		total += num
	}
	fmt.Printf("sum of all arugment is: %d\n", total)

	fmt.Println("first argument is: ", varg[0])
}

func strargv(argv ...string) {
	fmt.Println(argv)
}

func main() {

	core(1, 2)
	core(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4, 5}
	core(nums...)

	strargv("gaurav", "garg", "golang")
	strargv("garg", "gopher")
}
