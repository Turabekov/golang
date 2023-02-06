package main

import "fmt"

// 13.
// Type Student struct{
// 	Name string
// 	Grades []int
// 	Age int
// }
// Shu struct uchun getAverageGrade hamda getMaxGrade method larini yozing.

type Student struct {
	Name   string
	Grades []int
	Age    int
}

func (s Student) getAverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, val := range s.Grades {
		sum += val
	}
	return float64(sum / len(s.Grades))
}

func (s Student) getMaxGrade() int {

	if len(s.Grades) == 0 {
		return 0
	}

	max := 0

	for _, v := range s.Grades {
		if max <= v {
			max = v
		}
	}

	return max
}

func main() {
	studetn1 := Student{
		Name:   "Khumoyun",
		Grades: []int{5, 4, 3, 2, 1, 7},
		Age:    20,
	}

	fmt.Println(studetn1.getAverageGrade())
	fmt.Println(studetn1.getMaxGrade())
}
