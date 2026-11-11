package main

import "fmt"

// 接收者 (Receiver)
type Bulb struct{}

func NewBulb() *Bulb {
	return &Bulb{}
}

func (b *Bulb) TurnOn() {
	fmt.Println("Bulb has been lit")
}

func (b *Bulb) TurnOff() {
	fmt.Println("Darkness!")
}

// 命令 (Command)
type Command interface {
	Execute()
	Undo()
	Redo()
}

type TurnOn struct {
	bulb *Bulb
}

func NewTurnOn(bulb *Bulb) TurnOn {
	return TurnOn{bulb: bulb}
}

func (c *TurnOn) Execute() {
	c.bulb.TurnOn()
}

func (c *TurnOn) Undo() {
	c.bulb.TurnOff()
}

func (c *TurnOn) Redo() {
	c.Execute()
}

type TurnOff struct {
	bulb *Bulb
}

func NewTurnOff(bulb *Bulb) TurnOff {
	return TurnOff{bulb: bulb}
}

func (c *TurnOff) Execute() {
	c.bulb.TurnOff()
}

func (c *TurnOff) Undo() {
	c.bulb.TurnOn()
}

func (c *TurnOff) Redo() {
	c.Execute()
}

type RemoteControl struct{}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{}
}

func (rc *RemoteControl) Submit(command Command) {
	command.Execute()
}

func main() {
	bulb := NewBulb()
	turnOn := NewTurnOn(bulb)
	turnOff := NewTurnOff(bulb)

	remote := NewRemoteControl()
	remote.Submit(&turnOn)  // Bulb has been lit
	remote.Submit(&turnOff) // Darkness!
}
