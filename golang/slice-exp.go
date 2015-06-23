/* In slice if size of sub slice (in this example slice l) is smaller
 * then parent slice (in this example s) then address of sub slice will point
 * to the parent class. In this case any change in element of  sub slice will
 * effect to the parent slice. But if size of sub slice is grater then the
 * parent slice then address of sub slice will not point to the parent slice.
 * and any change in value of sub slice will not effect the parent slice.
 */

package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println ("Slice s is:", s)
	fmt.Println ("length of slice s is:", len(s))

	s[0] = "g"
	s[1] = "o"
	s[2] = "l"

	fmt.Println ("slice s is: ", s);

	l := s[1:3]

	fmt.Println ("slice l is:", l);
	
	l[0] = "a"
	fmt.Println ("slice l after changing l[0] to a is:", l);
	fmt.Println ("slice s is: ", s);

	l = append (l, "b")
	fmt.Println ("slice l after appending b to l is:", l);
	fmt.Println ("slice s is: ", s);

	fmt.Println ("Now size of parent slice s and sub slice l is same")

	fmt.Println ("Again modifying l[0] to R")
	l[0] = "R"
	fmt.Println ("slice l after changing l[0] to R is:", l);
	fmt.Println ("slice s is: ", s);
	fmt.Println ("parent and sub slice will not change. because size of sub slice is not smaller then size of parent slice.")

	fmt.Println ("after appending to slice l, size of slice l is greater then size of parent")
	l = append (l, "c")
	fmt.Println ("slice l after appending c to l is:", l);
	fmt.Println ("slice s is: ", s);

	l[0] = "G"
	fmt.Println ("slice l after changing l[0] to G is:", l);
	fmt.Println ("slice s is: ", s);

}
