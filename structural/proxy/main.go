package main

import "fmt"

type Door interface {
	Open(password string)
	Close()
}

type LabDoor struct{}

var _ Door = &LabDoor{}

func (d *LabDoor) Open(password string) {
	// Ignore the inputted password
	fmt.Println("Opening the lab door")
}

func (d *LabDoor) Close() {
	fmt.Println("Closing the lab door")
}

type SecuredLabDoor struct {
	labDoor *LabDoor
}

var _ Door = &SecuredLabDoor{}

func NewSecuredLabDoor() *SecuredLabDoor {
	return &SecuredLabDoor{labDoor: &LabDoor{}}
}

func (d *SecuredLabDoor) Open(password string) {
	if d.authenticate(password) {
		d.labDoor.Open("")
		return
	}
	fmt.Println("Big no! It ain't possible.")
}

func (d *SecuredLabDoor) Close() {
	d.labDoor.Close()
}

func (d *SecuredLabDoor) authenticate(password string) bool {
	return password == "$ecr@t"
}

func main() {
	securedLabDoor := NewSecuredLabDoor()
	securedLabDoor.Open("invalid") // Big no! It ain't possible.
	securedLabDoor.Open("$ecr@t")  // Opening the lab door
	securedLabDoor.Close()         // Closing the lab door
}
