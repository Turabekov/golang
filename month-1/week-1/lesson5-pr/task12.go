package main

import "fmt"

type MyHashSet struct {
	obj []int
}

func (this MyHashSet) Add(value int) MyHashSet {
	boolVar := false
	for _, v := range this.obj {
		if v != value {
			boolVar = true
		} else {
			boolVar = false
		}
	}
	if boolVar {
		this.obj = append(this.obj, value)
	}

	return this
}

func (this MyHashSet) Contains(value int) bool {
	boolVar := false
	for _, v := range this.obj {
		if v == value {
			boolVar = true
		} else {
			boolVar = false
		}
	}

	return boolVar
}

func (this MyHashSet) Remove(index int) MyHashSet {
	newSlice := []int{}

	for i, v := range this.obj {
		if i == index {
			continue
		}
		newSlice = append(newSlice, v)
	}

	this.obj = newSlice

	return this
}

func main() {
	newSet := MyHashSet{
		obj: []int{1, 2, 3, 4, 5, 6},
	}

	fmt.Println(newSet.Add(6))
	fmt.Println(newSet.Contains(7))
	fmt.Println(newSet.Remove(4))

}
