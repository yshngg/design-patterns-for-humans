package main

import "fmt"

type Animal interface {
	Accept(operation AnimalOperation)
}

type AnimalOperation interface {
	VisitMonkey(monkey *Monkey)
	VisitLion(lion *Lion)
	VisitDolphin(dolphin *Dolphin)
}

type Monkey struct{}

func (m *Monkey) Shout() {
	fmt.Println("Ooh oo aa aa!")
}

func (m *Monkey) Accept(operation AnimalOperation) {
	operation.VisitMonkey(m)
}

type Lion struct{}

func (l *Lion) Roar() {
	fmt.Println("Roaaar!")
}

func (l *Lion) Accept(operation AnimalOperation) {
	operation.VisitLion(l)
}

type Dolphin struct{}

func (d *Dolphin) Speak() {
	fmt.Println("Tuut tuttu tuutt!")
}

func (d *Dolphin) Accept(operation AnimalOperation) {
	operation.VisitDolphin(d)
}

type Speak struct{}

var _ AnimalOperation = &Speak{}

func (s *Speak) VisitMonkey(monkey *Monkey) {
	monkey.Shout()
}

func (s *Speak) VisitLion(lion *Lion) {
	lion.Roar()
}

func (s *Speak) VisitDolphin(dolphin *Dolphin) {
	dolphin.Speak()
}

type Jump struct{}

var _ AnimalOperation = &Jump{}

func (j *Jump) VisitMonkey(monkey *Monkey) {
	fmt.Println("Jumped 20 feet high! on to the tree!")
}

func (s *Jump) VisitLion(lion *Lion) {
	fmt.Println("Jumped 7 feet! Back on the ground!")
}

func (s *Jump) VisitDolphin(dolphin *Dolphin) {
	fmt.Println("Walked on water a little and disappeared")
}

func main() {
	monkey := &Monkey{}
	lion := &Lion{}
	dolphin := &Dolphin{}

	speak := &Speak{}
	monkey.Accept(speak)  // Ooh oo aa aa!
	lion.Accept(speak)    // Roaaar!
	dolphin.Accept(speak) // Tuut tuttu tuutt!

	jump := &Jump{}
	monkey.Accept(jump)  // Jumped 20 feet high! on to the tree!
	lion.Accept(jump)    // Jumped 7 feet! Back on the ground!
	dolphin.Accept(jump) // Walked on water a little and disappeared
}
