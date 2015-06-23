/* Channels: Channels are the pipes that connect concurrent goroutines. You
 * can send values into channels from one goroutine and receive those values
 * into another goroutine.
 * [1]. Create a new channel with make(chan val-type). Channels are typed by
 *      the values they convey.
 * [2]. Send a value into a channel using the channel <- syntax.
 * [3]. The <-channel syntax receives a value from the channel.
 * 
 * By default sends and receives block until both the sender and receiver are
 * ready. This property allowed us to wait at the end of our program for 
 * the "ping" message without having to use any other synchronization.
 * 
 * Use case of channel: If you want to received value from A function to 
 *                      another function with out returning from A function
 *			then create channel and send the value from A function
 *			to other function.
 */

package main

import (
 	"fmt"
)

func main() {
	/* [1] */
	messages := make (chan string)

	go func() {
		/* [2] */
		messages <- "ping"
	}()

	/*go func() {
		msg := <- messages
		fmt.Println (msg)
	}()*/
		

	/* [3] */
	msg := <- messages
	fmt.Println (msg)
}

