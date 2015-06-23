/* Worker pools in go: use to execute task concurrently.
 */

package main

import (
 	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing jobs", j)
		time.Sleep(time.Second)
		results <- j
	}
}


func main() {
	jobs := make (chan int, 100)
	results := make (chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close (jobs)

	for a := 1; a <=9; a++ {
		<- results
	}
}
