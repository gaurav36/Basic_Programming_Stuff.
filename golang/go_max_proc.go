

package main

import (
	"fmt"
	"runtime"
)

func Get_proc() {
	maxproc := runtime.GOMAXPROCS(0)
	fmt.Println ("max proc is: ", maxproc)
	
	max_cpu := runtime.NumCPU()

	fmt.Println ("max cpu is: ", max_cpu)
}

func main() {
	Get_proc()
}
