package main

import "fmt"

const (
	Pi = 3.14
)

type Circle struct {
	radi float64
}

func newcr(radi float64) Circle {
	return Circle{
		radi: radi,
	}

}
func main() {
	my_cr := newcr(11.5)
	my_cr.clac()

}
func (c Circle) clac() {
	circum := 2 * Pi * c.radi
	area := Pi * c.radi * c.radi
	fmt.Printf("%+v %+v", area, circum)
}
