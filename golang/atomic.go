/* atomic in Go: to perform atomic operation in go language
 * Go having "sync/atomic" package for performing atomic operation
 * [1]. runtime.Gosched(): used for Allow other goroutines to proceed.
 *      for more reference about it:
 *      http://stackoverflow.com/questions/13107958/what-exactly-does-runtime-gosched-do
 */

package main

import (
 	"fmt"
	"time"
	"runtime"
	"sync/atomic"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println ("ops is: ", ops)
	fmt.Println ("opsFinal is: ", opsFinal)
}
 	
