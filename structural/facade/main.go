package main

import "fmt"

type Computer struct{}

func (c *Computer) GetElectricShock() {
	fmt.Println("Ouch!")
}

func (c *Computer) MakeSound() {
	fmt.Println("Beep beep!")
}

func (c *Computer) ShowLoadingScreen() {
	fmt.Println("Loading..")
}

func (c *Computer) BAM() {
	fmt.Println("Ready to be used!")
}

func (c *Computer) CloseEverything() {
	fmt.Println("Bup bup bup buzzzz!")
}

func (c *Computer) Sooth() {
	fmt.Println("Zzzzz")
}

func (c *Computer) PullCurrent() {
	fmt.Println("Haaah!")
}

type ComputerFacade struct {
	computer *Computer
}

func (c ComputerFacade) TurnOn() {
	c.computer.GetElectricShock()
	c.computer.MakeSound()
	c.computer.ShowLoadingScreen()
	c.computer.BAM()
}

func (c ComputerFacade) TurnOff() {
	c.computer.CloseEverything()
	c.computer.PullCurrent()
	c.computer.Sooth()
}

func main() {
	computer := ComputerFacade{computer: &Computer{}}
	computer.TurnOn()  // Ouch! Beep beep! Loading.. Ready to be used!
	computer.TurnOff() // Bup bup bup buzzzz! Haaah! Zzzzz
}
