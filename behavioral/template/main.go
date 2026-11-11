package main

import "fmt"

type Builder interface {
	test()
	lint()
	assemble()
	deploy()
}

type builder struct {
	Builder
}

func (b *builder) build() {
	b.test()
	b.lint()
	b.assemble()
	b.deploy()
}

type AndroidBuilder struct{}

var _ Builder = &AndroidBuilder{}

func (ab *AndroidBuilder) test() {
	fmt.Println("Running android tests")
}

func (ab *AndroidBuilder) lint() {
	fmt.Println("Linting the android code")
}

func (ab *AndroidBuilder) assemble() {
	fmt.Println("Assembling the android build")
}

func (ab *AndroidBuilder) deploy() {
	fmt.Println("Deploying android build to server")
}

type IosBuilder struct{}

var _ Builder = &IosBuilder{}

func (ab *IosBuilder) test() {
	fmt.Println("Running ios tests")
}

func (ab *IosBuilder) lint() {
	fmt.Println("Linting the ios code")
}

func (ab *IosBuilder) assemble() {
	fmt.Println("Assembling the ios build")
}

func (ab *IosBuilder) deploy() {
	fmt.Println("Deploying ios build to server")
}

func main() {
	androidBuilder := &AndroidBuilder{}
	b := builder{
		Builder: androidBuilder,
	}
	b.build()
	// 输出:
	// Running android tests
	// Linting the android code
	// Assembling the android build
	// Deploying android build to server

	iosBuilder := &IosBuilder{}
	b = builder{
		Builder: iosBuilder,
	}
	b.build()
	// 输出:
	// Running ios tests
	// Linting the ios code
	// Assembling the ios build
	// Deploying ios build to server
}
