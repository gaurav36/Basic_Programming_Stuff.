/* Nested structure programme in go
 */

package main

import (
 	"fmt"
)


type base1 struct {
	name string
	id   int
	tmp *base1
}

type base2 struct {
	cname string
	temp  base1
}

func main () {
	var st base2
	st.cname = "gaurav"
	st.temp.name = "garg"
	st.temp.id = 10

	fmt.Println ("base2 cname is: ", st.cname)
	fmt.Println ("base1 name is: ", st.temp.name)
	fmt.Println ("base1 ID is: ", st.temp.id)
}
