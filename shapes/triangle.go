package main

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 1.5 * t.base * t.height
}
