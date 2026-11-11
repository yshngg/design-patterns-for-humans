package main

import "fmt"

type KarakTea struct{}

type TeaMaker struct {
	availableTea map[string]*KarakTea
}

func NewTeaMaker() *TeaMaker {
	return &TeaMaker{
		availableTea: make(map[string]*KarakTea, 0),
	}
}

func (m *TeaMaker) Make(preference string) *KarakTea {
	if m.availableTea[preference] == nil {
		m.availableTea[preference] = &KarakTea{}
	}

	return m.availableTea[preference]
}

type TeaShop struct {
	orders   map[int]*KarakTea
	teaMaker *TeaMaker
}

func NewTeaShop(teaMaker *TeaMaker) *TeaShop {
	return &TeaShop{
		orders:   make(map[int]*KarakTea, 0),
		teaMaker: teaMaker,
	}
}

func (s *TeaShop) TakeOrder(teaType string, table int) {
	s.orders[table] = s.teaMaker.Make(teaType)
}

func (s *TeaShop) Serve() {
	for table, tea := range s.orders {
		fmt.Printf("Serving tea %s to table #%d\n", tea, table)
	}
}

func main() {
	teaMaker := NewTeaMaker()
	shop := NewTeaShop(teaMaker)

	shop.TakeOrder("less sugar", 1)
	shop.TakeOrder("more milk", 2)
	shop.TakeOrder("without sugar", 5)

	shop.Serve()
	// Serving tea &{} to table #1
	// Serving tea &{} to table #2
	// Serving tea &{} to table #5
}
