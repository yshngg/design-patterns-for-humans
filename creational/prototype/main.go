package main

import "fmt"

type Sheep struct {
	name     string
	category string
}

func NewSheep(name, category string) *Sheep {
	return &Sheep{
		name:     name,
		category: category,
	}
}

func (s Sheep) GetName() string {
	return s.name
}

func (s Sheep) GetCategory() string {
	return s.category
}

func (s *Sheep) SetName(name string) {
	s.name = name
}

func (s *Sheep) SetCategory(category string) {
	s.category = category
}

func (s Sheep) Clone() Sheep {
	return Sheep{
		name:     s.name,
		category: s.category,
	}
}

func main() {
	original := NewSheep("Jolly", "Mountain Sheep")
	fmt.Println(original.GetName())     // Jolly
	fmt.Println(original.GetCategory()) // Mountain Sheep

	cloned := original.Clone()
	cloned.SetName("Dolly")
	fmt.Println(cloned.GetName())     // Dolly
	fmt.Println(cloned.GetCategory()) // Mountain Sheep
}
