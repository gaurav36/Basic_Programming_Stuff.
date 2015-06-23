/* link list by using container package 
 */

package main

import (
	"fmt"
	"container/list"
)

func Print_list (l *list.List){
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println (e.Value)
	}
}

type Data struct {
	Name string
	Age  int
}


func main() {
	l := list.New()

	e4 := l.PushBack(4)

	e5 := l.PushFront(5)

	l.InsertBefore (6, e4)
	l.InsertAfter (7, e5)
	l.PushBack (Data{Name: "Gaurav", Age: 23})
	l.InsertBefore (Data{Name: "Garg", Age: 23}, e4)
	
	f := l.Front()
	fmt.Println ("front element is: ", f.Value)
	
	b := l.Back()
	fmt.Println ("back element is: ", b.Value)

	Print_list (l)

	fmt.Println ("Length of link list is: ", l.Len()) /* Print the length of list*/
}

