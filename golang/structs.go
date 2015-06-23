/* Structure in go: It works as same logic as structure in C/C++ programming language work
 */

package main

import (
	"fmt"
)

type employee struct {
	name string
	id   int
}

func main() {
	fmt.Println (employee{"Bob", 20})

	fmt.Println (employee{name: "Gaurav", id: 36})

	fmt.Println (employee{name: "Garg"})
	
	fmt.Println (&employee{name: "Adams", id: 32})

	s := employee{name: "john", id: 12}

	fmt.Println ("name is: ", s.name, "\nid is: ", s.id)

	sp := &s

	fmt.Println (sp.name) /* here sp->name will not work even if it is pointer*/
	fmt.Println (sp.id)
}
