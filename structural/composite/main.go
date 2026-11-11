package main

import (
	"fmt"
	"strings"
)

type Employee interface {
	GetName() string
	SetSalary(salary float32)
	GetSalary() float32
	GetRoles() []string
}

type Developer struct {
	name   string
	salary float32
	roles  []string
}

var _ Employee = &Developer{}

func NewDeveloper(name string, salary float32) *Developer {
	return &Developer{name: name, salary: salary}
}

func (d Developer) GetName() string {
	return d.name
}

func (d *Developer) SetSalary(salary float32) {
	d.salary = salary
}

func (d Developer) GetSalary() float32 {
	return d.salary
}

func (d Developer) GetRoles() []string {
	return d.roles
}

type Designer struct {
	name   string
	salary float32
	roles  []string
}

var _ Employee = &Designer{}

func NewDesigner(name string, salary float32) Designer {
	return Designer{name: name, salary: salary}
}

func (d Designer) GetName() string {
	return d.name
}

func (d *Designer) SetSalary(salary float32) {
	d.salary = salary
}

func (d Designer) GetSalary() float32 {
	return d.salary
}

func (d Designer) GetRoles() []string {
	return d.roles
}

type Organization struct {
	employees []Employee
}

func (o *Organization) AddEmployee(employee Employee) {
	o.employees = append(o.employees, employee)
}

func (o Organization) GetNetSalaries() float32 {
	netSalary := float32(0)
	for _, employee := range o.employees {
		netSalary += employee.GetSalary()
	}
	return netSalary
}

// func main() {
// 	// 准备员工
// 	john := NewDeveloper("John Doe", 12000)
// 	jane := NewDeveloper("Jane Doe", 15000)

// 	// 把他们加入组织
// 	organization := Organization{}
// 	organization.AddEmployee(john)
// 	organization.AddEmployee(jane)

// 	fmt.Println("Net salaries:", organization.GetNetSalaries()) // Net salaries: 27000
// }

type Component interface {
	Find(text string) bool
}

type File struct {
	name    string
	content string
}

var _ Component = &File{}

func NewFile(name, content string) *File {
	return &File{
		name:    name,
		content: content,
	}
}

func (f *File) Find(text string) bool {
	return strings.Contains(f.content, text)
}

func (f File) Name() string {
	return f.name
}

type Folder struct {
	name       string
	components []Component
}

var _ Component = &Folder{}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f *Folder) Find(text string) bool {
	for _, component := range f.components {
		if !component.Find(text) {
			continue
		}
		return true
	}
	return false
}

func (f Folder) Name() string {
	return f.name
}

func (f *Folder) AddFile(file *File) {
	if file == nil {
		return
	}
	f.components = append(f.components, file)
}

func (f *Folder) AddFolder(folder *Folder) {
	if folder == nil {
		return
	}
	f.components = append(f.components, folder)
}

func main() {
	file1 := NewFile("file1", "some text")
	file2 := NewFile("file2", "some text also")

	folder := NewFolder("dirname")
	folder.AddFile(file1)
	folder.AddFile(file2)

	fmt.Println(folder.Find("also")) // true
}
