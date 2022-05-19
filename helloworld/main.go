package main

import (
	"fmt"
)

// func fibonacci(n int) []int {
// 	if n < 2 {
// 		return make([]int, 0)
// 	}

// 	nums := make([]int, n)
// 	nums[0], nums[1] = 1, 1

// 	for i := 2; i < n; i++ {
// 		nums[i] = nums[i-1] + nums[i-2]
// 	}

// 	return nums
// }

func main() {
	var s Shape = Square{3}
	fmt.Println("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

// type triangel struct {
// 	size int
// }
// type coloredTriangle struct {
// 	triangel
// 	color string
// }

// func (t triangel) perimeter() int {
// 	return t.size * 3
// }

// func (t *triangel) doubleSize() {
// 	t.size *= 2
// }

// func (t coloredTriangle) perimeter() int {
// 	return t.triangel.perimeter()
// }

// func (t coloredTriangle) perimeter() int {
// 	return t.size * 3 * 2
// }

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}
func (s Square) Perimeter() float64 {
	return s.size * 4
}
