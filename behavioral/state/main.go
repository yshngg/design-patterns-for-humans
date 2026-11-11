package main

import "errors"

type PhoneState interface {
	PickUp() (PhoneState, error)
	HangUp() (PhoneState, error)
	Dial() (PhoneState, error)
}

type PhoneStateIdle struct{}

var _ PhoneState = &PhoneStateIdle{}

func (ps *PhoneStateIdle) PickUp() (PhoneState, error) {
	return &PhoneStatePickedUp{}, nil
}

func (ps *PhoneStateIdle) HangUp() (PhoneState, error) {
	return nil, errors.New("already idle")
}

func (ps *PhoneStateIdle) Dial() (PhoneState, error) {
	return nil, errors.New("unable to dial in idle state")
}

type PhoneStatePickedUp struct{}

var _ PhoneState = &PhoneStatePickedUp{}

func (ps *PhoneStatePickedUp) PickUp() (PhoneState, error) {
	return nil, errors.New("already picked up")
}

func (ps *PhoneStatePickedUp) HangUp() (PhoneState, error) {
	return &PhoneStateIdle{}, nil
}

func (ps *PhoneStatePickedUp) Dial() (PhoneState, error) {
	return &PhoneStateCalling{}, nil
}

type PhoneStateCalling struct{}

var _ PhoneState = &PhoneStateCalling{}

func (ps *PhoneStateCalling) PickUp() (PhoneState, error) {
	return nil, errors.New("already picked up")
}

func (ps *PhoneStateCalling) HangUp() (PhoneState, error) {
	return &PhoneStateIdle{}, nil
}

func (ps *PhoneStateCalling) Dial() (PhoneState, error) {
	return nil, errors.New("already dialing")
}

type Phone struct {
	state PhoneState
}

func NewPhone() *Phone {
	return &Phone{
		state: &PhoneStateIdle{},
	}
}

func (p *Phone) PickUp() (err error) {
	p.state, err = p.state.PickUp()
	return
}

func (p *Phone) HangUP() (err error) {
	p.state, err = p.state.HangUp()
	return
}

func (p *Phone) Dial() (err error) {
	p.state, err = p.state.Dial()
	return
}

func main() {
	phone := NewPhone()
	phone.PickUp()
	phone.Dial()
}
