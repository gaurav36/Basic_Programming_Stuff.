/* Map: hashes or dicts in other language
 * [1]. To create an empty map, use the builtin make: make(map[key-type]val-type).
 * [2]. Set key/value pairs using typical name[key] = val syntax.
 * [3]. The optional second return value when getting a value from a map
 *      indicates if the key was present in the map. This can be used to
 *      disambiguate between missing keys and keys with zero values like 0
 *      or "". Here we didn’t need the value itself, so we ignored it with
 *      the blank identifier _,
 * [4]. You can also declare and initialize a new map in the same line with
 *      this syntax.
 */

package main

import (
 	"fmt"
)

func main() {
        /* [1] */
	m := make(map[string]int)

	/* [2] */
	m["key1"] = 7
	m["key2"] = 12

	fmt.Println ("map m is: ", m)

	/* Get a value for a key with name[key]. */
	v1 := m["key1"]
	v2 := m["key2"]
	fmt.Println ("v1 is: ", v1,"\n", "v2 is: ", v2)

	/* Printing number of key/value pair*/
	fmt.Println("len:", len(m))
	
	/* remove the key value pair from map */
	delete (m, "key1")
	fmt.Println ("map after deleting key1: ", m)

	/* [3] */
	/* prs1 will have first return value, prs2 will have second return value*/
	prs1 := m["key1"]
	_, prs2 := m["key1"] /* Ignore first return value. give 2nd return value
	                      * here ignoring is done by "_"
			      */
	fmt.Println ("parsing of deleted key1: ", prs1)
	fmt.Println ("parsing of deleted key1: ", prs2)

        /* updating value in map*/
	m["key2"] = 0
	   prs1 = m["key2"]  /* don't need to use ":="  because prs1 already declared in above line's*/
	_, prs2 = m["key2"] /* Ignore first return value. give 2nd return value*/
	fmt.Println ("parsing of deleted key1: ", prs1)
	fmt.Println ("parsing of deleted key1: ", prs2)

	/* Printing map*/
	fmt.Println ("after updating value in map m is: ", m)

	/* maps are dynamically sized. Appending more key in map*/
	m["key3"] = 4
	m["key4"] = 5
	fmt.Println ("map m is: ", m)

	prs3, prs4 := m["key4"] /*don't ignore any value from map*/
	fmt.Println ("with out ignoring","prs 3 is: ", prs3, " prs 4 is: ", prs4)

	/* [4] */
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println ("new dynamic map n is: ", n)



}
