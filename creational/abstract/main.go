package main

import "fmt"

type Door interface {
	Open()
	Close()
}

type WoodenDoor struct{}

var _ Door = &WoodenDoor{}

func (w *WoodenDoor) Open() {
	fmt.Println("Opening the wooden door")
}

func (w *WoodenDoor) Close() {
	fmt.Println("Closing the wooden door")
}

type IronDoor struct{}

var _ Door = &IronDoor{}

func (w *IronDoor) Open() {
	fmt.Println("Opening the iron door")
}

func (w *IronDoor) Close() {
	fmt.Println("Closing the iron door")
}

type DoorFittingExpert interface {
	Fit()
}

type Carpenter struct{}

var _ DoorFittingExpert = &Carpenter{}

func (c *Carpenter) Fit() {
	fmt.Println("Fitting wooden doors")
}

type Welder struct{}

var _ DoorFittingExpert = &Welder{}

func (e *Welder) Fit() {
	fmt.Println("Fitting iron doors")
}

// 木门工厂，返回木匠和木门
type DoorFactory interface {
	MakeDoor() Door
	MakeFittingExpert() DoorFittingExpert
}

type WoodenDoorFactory struct{}

func (f *WoodenDoorFactory) MakeDoor() Door {
	return &WoodenDoor{}
}

func (f *WoodenDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return &Carpenter{}
}

// 铁门工厂，返回铁门和对应的安装专家
type IronDoorFactory struct{}

func (f *IronDoorFactory) MakeDoor() Door {
	return &IronDoor{}
}

func (f *IronDoorFactory) MakeFittingExpert() DoorFittingExpert {
	return &Welder{}
}

func main() {
	woodenFactory := WoodenDoorFactory{}
	door := woodenFactory.MakeDoor()
	expert := woodenFactory.MakeFittingExpert()

	expert.Fit() // 输出: Fitting wooden doors
	door.Open()  // 输出: Opening the wooden door

	ironFactory := IronDoorFactory{}
	door = ironFactory.MakeDoor()
	expert = ironFactory.MakeFittingExpert()

	expert.Fit() // 输出: Fitting iron doors
	door.Open()  // 输出: Opening the iron door
}
