/* Switch case in go language:
 * [1]. There is no need to use break in b/w switch case
 * [2]. You can use commas to separate multiple expressions in the same case
 *      statement. We use the optional default case in this example as well.
 * [3]. switch without an expression is an alternate way to express if/else logic.
 *      Here we also show how the case expressions can be non-constants.
 */


package main

import (
	"fmt"
	"time"
)

func main (){
	i := 2

	fmt.Print ("write ", i, " as ")

        /* [1] */
	switch i {
	case 1:
	        fmt.Println ("One")
	case 2:
	        fmt.Println ("Two")
		/* break;  [1] its optional to use break statement here.*/
	case 3:
		fmt.Println ("Three")
	}

        /* [2] */
	switch time.Now().Weekday () {
	case time.Saturday, time.Sunday:
		fmt.Println ("It's a weekend")
	default:
		fmt.Println ("It's a weekday");
	}

	/* [3] */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println ("Actual time is", t)
		fmt.Println ("It's before noon")
	default:
		fmt.Println ("Actual time is", t)
		fmt.Println ("It's after noon")
	}
}
