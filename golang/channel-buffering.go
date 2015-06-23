/* Channel buffering in Go:
 *
 */

package main

import (
	"fmt"
)

func main() {

	/* Here we make a channel of strings buffering up to 2 values. */
	messages := make(chan string, 3)

	messages <- "buffered"
	messages <- "channels"
	messages <- "golang"

      /*msg1 := <- messages
	msg2 := <- messages
	msg3 := <- messages

	fmt.Println (msg1)
	fmt.Println (msg2)
	fmt.Println (msg3)*/


	fmt.Println (<- messages)
	fmt.Println (<- messages)
	fmt.Println (<- messages)
	//fmt.Println (<- messages) /* we have already received max limit 3 value
				       it will throw error if we receive more channel value*/
}

