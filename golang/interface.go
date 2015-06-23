/* Interface in Go: Interface named as collection of method signature.
 * 
 */

package main

import (
 	"fmt"
	"math"
)


/*type geometry interface {
	area() float64
	perim() float64
}*/

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area () float64 {
	return r.width * r.height
}

func (r rect) perim () float64 {
	return 2*r.width + 2*r.height
}

func (r circle) area () float64 {
	return math.Pi * r.radius * r.radius
}

func (r circle) perim () float64 {
	return 2 * math.Pi * r.radius
}

/*
func measure(g geometry) {
	fmt.Println (g)
	fmt.Println (g.area())
	fmt.Println (g.perim())
}*/

func main() {
	
	r := rect {width: 10, height: 5}
	c := circle {radius: 3}

	//measure(r)
	//measure (c)

	fmt.Println(r.area ());
	fmt.Println(c.area ());
}
