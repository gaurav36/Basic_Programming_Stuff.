/* Channel Synchronization in Go: use channels to synchronize execution to
 * blocking receive to wait for a goroutine to finish.	
 */

package main

import (
 	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println ("Working...")
	time.Sleep (time.Second)
	fmt.Println ("done")
	
	/* Send a value to notify that we’re done. */
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)
	
	/* Block until we receive a notification from the worker on the channel.*/
	<- done


	/* If you removed the <- done line from this program, the program
	 * would exit before the worker even started.
	 */
}
