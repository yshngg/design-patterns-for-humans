package main

import (
	"fmt"
	"reflect"
	"sync"
)

type President struct{}

var (
	once sync.Once

	instance *President
)

func NewPresident() *President {
	once.Do(func() {
		instance = new(President)
	})

	return instance
}

func main() {
	president1 := NewPresident()
	president2 := NewPresident()
	fmt.Println(reflect.DeepEqual(president1, president2)) // true
}
