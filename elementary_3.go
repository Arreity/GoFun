package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What is your name?")
	scanner.Scan()
	name := scanner.Text()

	if name == "Alice" || name == "Bob" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Go Away!")
	}
}
