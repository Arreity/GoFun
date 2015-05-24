package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Mister, What is your name?")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello", name)
}
