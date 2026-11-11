package main

type Lion interface {
	Roar()
}

type AfricanLion struct{}

var _ Lion = &AfricanLion{}

func (l *AfricanLion) Roar() {}

type AsianLion struct{}

var _ Lion = &AsianLion{}

func (l *AsianLion) Roar() {}

type Hunter struct{}

func (h *Hunter) Hunt(lion Lion) {
	lion.Roar()
}

type WildDog struct{}

func (d *WildDog) Bark() {}

type WildDogAdapter struct {
	Dog WildDog
}

var _ Lion = &WildDogAdapter{}

func (d *WildDogAdapter) Roar() {
	d.Dog.Bark()
}

func main() {
	wildDog := WildDog{}
	wildDogAdapter := WildDogAdapter{Dog: wildDog}

	hunter := Hunter{}
	hunter.Hunt(&wildDogAdapter)
}
