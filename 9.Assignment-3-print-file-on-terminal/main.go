package main

import (
	"io"
	"log"
	"os"
)

// run the command
// option 1: 1st: create a file like my_text.txt, 2nd: go run main.go <<filename>>.txt
// option 2: if you want to read main.go, because main.go is a sensitive case for go compiler, so
//           you should build the main.go first, then run the command "./main main.go"
func main() {
	fileName := os.Args[1]

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// manually read file and print its content on the terminal
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// use io.Copy to read file and prints its content on terminal
	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err)
	}
}
