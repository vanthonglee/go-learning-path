package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type vietnameseBot struct{}

func main() {
	eb := englishBot{}
	vb := vietnameseBot{}

	printGreeting(eb)
	printGreeting(vb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There"
}

func (vb vietnameseBot) getGreeting() string {
	// Very custom logic for generating an Vietnamese greeting
	return "Xin Chao"
}
