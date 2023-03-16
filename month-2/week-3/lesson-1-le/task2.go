package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Student struct {
	Name  string
	Speed int
}

func (c Student) GetSpeed() time.Duration {
	return time.Duration(rand.Intn(1000))
}

func Run(student Student, idx int, empty chan struct{}) {
	speed := student.GetSpeed()

	fmt.Printf("%d. %s imtihon boshlandi \n", idx, student.Name)

	time.Sleep(time.Millisecond * speed)
	fmt.Printf("%d. %s imtihon topshirdi \n", idx, student.Name)

	empty <- struct{}{}
}

func main() {
	// var students = []Student{
	// 	{Name: "Xumoyun"},
	// 	{Name: "Baxrom"},
	// 	{Name: "Shaxzod"},
	// 	{Name: "Mirazam"},
	// 	{Name: "Azam"},
	// }

	// empty := make(chan struct{})

	// fmt.Println("Started")
	// for idx, st := range students {
	// 	go Run(st, idx+1, empty)
	// }
	// <-empty

	// fmt.Scanln()

	fmt.Println(runtime.NumCPU())
}

// Korzinka

// if math.Abs(float64(val[i]) - float64(val[j])) <= float64(k) {
//     return true
// }
