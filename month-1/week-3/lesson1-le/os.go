package main

import "fmt"

type Os struct {
	Name string
}

type MacOs struct {
	Os
}

type Linux struct {
	Os
}

type Windows struct {
	Os
}

func (o Os) GetName() string {
	return o.Name
}

// overrading
func (m MacOs) GetName() string {
	// Super call
	// m.Os.GetName()
	return m.Os.GetName() + " M2"
}
func (w Windows) GetName() string {
	return w.Os.GetName() + "Max"
}

type OsInterface interface {
	GetName() string
}

func Run(os OsInterface) {
	fmt.Println("Loading " + os.GetName() + "...")
}

func main() {
	Run(MacOs{Os{Name: "Macintosh"}})
	Run(Linux{Os{Name: "Fedora 37"}})
	Run(Windows{Os{Name: "Windows 11 pro"}})

}
