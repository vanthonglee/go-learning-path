package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[int]string)

	// colors["white"] = "#ffffff"
	// Cannot use this syntax to access white like struts
	// structName.white
	// instead, colors["white"] becuase all keys inside of our map are typed

	// colors[10] = "ten"

	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
