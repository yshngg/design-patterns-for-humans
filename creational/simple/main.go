package main

import "fmt"

type Door interface {
	GetWidth() float32
	GetHeight() float32
}

type WoodenDoor struct {
	width  float32
	height float32
}

var _ Door = &WoodenDoor{}

func (d WoodenDoor) GetWidth() float32 {
	return d.width
}

func (d WoodenDoor) GetHeight() float32 {
	return d.height
}

func MakeDoor(weight, height float32) Door {
	return WoodenDoor{
		width:  weight,
		height: height,
	}
}

func main() {
	// 给我造一扇 100x200 的门
	door := MakeDoor(100, 200)
	fmt.Println("Width:", door.GetWidth())
	fmt.Println("Height:", door.GetHeight())

	// 给我造一扇 50x100 的门
	door = MakeDoor(50, 100)
	fmt.Println("Width:", door.GetWidth())
	fmt.Println("Height:", door.GetHeight())
}
