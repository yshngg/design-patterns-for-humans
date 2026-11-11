package main

import "fmt"

type Theme interface {
	GetColor() string
}

type DarkTheme struct{}

var _ Theme = &DarkTheme{}

func (t *DarkTheme) GetColor() string {
	return "Dark Black"
}

type LightTheme struct{}

var _ Theme = &LightTheme{}

func (t *LightTheme) GetColor() string {
	return "Off White"
}

type AquaTheme struct{}

var _ Theme = &AquaTheme{}

func (t *AquaTheme) GetColor() string {
	return "Light Blue"
}

type WebPage interface {
	GetContent() string
}

type About struct {
	Theme
}

var _ WebPage = &About{}

func (a *About) GetContent() string {
	return "About page in " + a.Theme.GetColor()
}

type Careers struct {
	Theme
}

var _ WebPage = &Careers{}

func (c *Careers) GetContent() string {
	return "Careers page in " + c.Theme.GetColor()
}

func main() {
	darkTheme := DarkTheme{}

	about := About{Theme: &darkTheme}
	careers := Careers{Theme: &darkTheme}

	fmt.Println(about.GetContent())   // About page in Dark Black
	fmt.Println(careers.GetContent()) // Careers page in Dark Black
}
