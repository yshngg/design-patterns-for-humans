package main

import "fmt"

type Burger struct {
	size                               int
	cheese, pepperoni, lettuce, tomato bool
}

type BurgerBuilder struct {
	size                               int
	cheese, pepperoni, lettuce, tomato bool
}

func NewBurgerBuilder(size int) *BurgerBuilder {
	return &BurgerBuilder{
		size: size,
	}
}

func (b *BurgerBuilder) AddCheese() *BurgerBuilder {
	b.cheese = true
	return b
}

func (b *BurgerBuilder) AddPepperoni() *BurgerBuilder {
	b.pepperoni = true
	return b
}

func (b *BurgerBuilder) AddLettuce() *BurgerBuilder {
	b.lettuce = true
	return b
}

func (b *BurgerBuilder) AddTomato() *BurgerBuilder {
	b.tomato = true
	return b
}

func (b *BurgerBuilder) Build() Burger {
	return Burger{
		size:      b.size,
		cheese:    b.cheese,
		pepperoni: b.pepperoni,
		lettuce:   b.lettuce,
		tomato:    b.tomato,
	}
}

func NewBurger(size int, cheese, pepperoni, tomato, lettuce bool) *Burger {
	return &Burger{}
}

func main() {
	burger := NewBurgerBuilder(14).
		AddCheese().
		AddPepperoni().
		AddLettuce().
		AddTomato().
		Build()
	fmt.Println(burger) // 输出: {14 true true true true}
}
