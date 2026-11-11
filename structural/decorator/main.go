package main

import "fmt"

type Coffee interface {
	GetCost() int
	GetDescription() string
}

type SimpleCoffee struct{}

var _ Coffee = SimpleCoffee{}

func (c SimpleCoffee) GetCost() int {
	return 10
}

func (c SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}

type MilkCoffee struct {
	coffee Coffee
}

var _ Coffee = MilkCoffee{}

func (c MilkCoffee) GetCost() int {
	return c.coffee.GetCost() + 2
}

func (c MilkCoffee) GetDescription() string {
	return c.coffee.GetDescription() + ", milk"
}

type WhipCoffee struct {
	coffee Coffee
}

var _ Coffee = WhipCoffee{}

func (c WhipCoffee) GetCost() int {
	return c.coffee.GetCost() + 5
}

func (c WhipCoffee) GetDescription() string {
	return c.coffee.GetDescription() + ", whip"
}

type VanillaCoffee struct {
	coffee Coffee
}

var _ Coffee = VanillaCoffee{}

func (c VanillaCoffee) GetCost() int {
	return c.coffee.GetCost() + 3
}

func (c VanillaCoffee) GetDescription() string {
	return c.coffee.GetDescription() + ", vanilla"
}

func main() {
	var someCoffee Coffee
	someCoffee = SimpleCoffee{}
	fmt.Println(someCoffee.GetCost())        // 10
	fmt.Println(someCoffee.GetDescription()) // Simple coffee

	someCoffee = MilkCoffee{coffee: someCoffee}
	fmt.Println(someCoffee.GetCost())        // 12
	fmt.Println(someCoffee.GetDescription()) // Simple coffee, milk

	someCoffee = WhipCoffee{coffee: someCoffee}
	fmt.Println(someCoffee.GetCost())        // 17
	fmt.Println(someCoffee.GetDescription()) // Simple coffee, milk, whip

	someCoffee = VanillaCoffee{coffee: someCoffee}
	fmt.Println(someCoffee.GetCost())        // 20
	fmt.Println(someCoffee.GetDescription()) // Simple coffee, milk, whip, vanilla
}
