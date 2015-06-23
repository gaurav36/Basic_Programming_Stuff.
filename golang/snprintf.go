

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	temp := fmt.Sprintf ("my temp using Sprintf is %s", "golang")

	fmt.Println ("temp is: ", temp)

	fmt.Println ("size of temp is: ", unsafe.Sizeof(temp))

}
