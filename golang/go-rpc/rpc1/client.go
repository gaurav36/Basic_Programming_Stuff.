package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

func main() {
	//  Dial connects to an RPC server at the specified network address.
	client, err := rpc.Dial("tcp", "localhost:42585")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply bool

		// Call invokes the named function, waits for it to complete,
		// and returns its error status.
		err = client.Call("Listener.Printline", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}
}
