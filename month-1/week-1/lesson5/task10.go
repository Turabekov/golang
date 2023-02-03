package main

import "fmt"

// 10.  Rectangle va Circle struct yarating hamda ularni yuzlari va perimetrlani topuvchi metodlar yozing.

const pi = 3.14

type Rectangle struct {
	a int
	b int
}

type Circle struct {
	r float64
}

func (r Rectangle) calculateAreaAndPerimetr() (int, int) {
	return r.a * r.b, 2 * (r.a + r.b)
}

func (c Circle) calcCircleAreaAndPerimetr() (float64, float64) {
	return c.r * pi, 2 * pi * c.r
}

func main() {

	rectangle1 := Rectangle{10, 20}
	circle1 := Circle{20}

	s2, p2 := circle1.calcCircleAreaAndPerimetr()

	fmt.Println(rectangle1.calculateAreaAndPerimetr())
	fmt.Printf("%.2f %.2f", s2, p2)
}
