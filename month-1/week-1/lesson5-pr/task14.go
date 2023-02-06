package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

// Engineer uchun Print() va GetProjectPriority() method larini yarating.

func (e Engineer) Print() {
	fmt.Println("Name", e.Name, "Age:", e.Age, "Project:", e.Project.Name)
}

func (e Engineer) GetProjectPriority() string {
	return e.Project.Priority
}

func main() {
	engineer1 := Engineer{
		Name: "Khumoyun",
		Age:  20,
		Project: Project{
			Name:         "Think a lot",
			Priority:     "active",
			Technologies: []string{"JavaScript", "Golang", "Docker"},
		},
	}

	engineer1.Print()
	fmt.Println(engineer1.GetProjectPriority())

}
