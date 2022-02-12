package main

import "fmt"

func main() {

	// var colors map[string]string

	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// }

	// colors["white"] = "#ffffff"
	// Cannot use this syntax to access white like struts
	// structName.white
	// instead, colors["white"] becuase all keys inside of our map are typed

	colors[10] = "ten"

	delete(colors, 10)

	fmt.Println(colors)
}
