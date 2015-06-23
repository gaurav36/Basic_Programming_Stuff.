/* Wait group: is one of the feature in Golang for goroutine for waiting until
 * all goroutines have finished.
 *
 * A WaitGroup waits for a collection of goroutines to finish. The main
 * goroutine calls Add to set the number of goroutines to wait for. Then
 * each of the goroutines runs and calls Done when finished. At the same
 * time, Wait can be used to block until all goroutines have finished.
 *
 * To use sync.WaitGroup we:
 * 1. Create a new instance of a sync.WaitGroup (we’ll call it wg)
 * 2. Call wg.Add(n) where n is the number of goroutines to wait for
 *    (we can also call wg.Add(1) n times)
 * 3. Execute defer wg.Done() in each goroutine to indicate that goroutine
 *    is finished executing to the WaitGroup (see defer)
 * 4. Call wg.Wait() where we want to block.
 * 
 * NOTE:
 * Channel can also be used for goroutine for waiting until all goroutines
 * have finished. it introduces some additional weirdness in the case that
 * we don’t actually know how many goroutines we want to spin up ahead of time.
 */


package main


import (
	"fmt"
	"sync"
)

var k int

func Prime1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println ("k is in prime1  ",k)
}

func Prime2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println ("k is in prime2 ",k)

}

func main () {
	var wg sync.WaitGroup
	//wg := &sync.WaitGroup{}

	fmt.Println ("k is in main", k)

	wg.Add(1)
	go Prime1(&wg)
	//go Prime1(wg)

	wg.Add(1)
	go Prime2(&wg)  /* call by reference */
	//go Prime2(wg) /* call by value*/

	wg.Wait()
}
