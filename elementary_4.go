package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What number do you want the sum of?")
	scanner.Scan()
	input := scanner.Text()
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	fmt.Println("The sum is", sum)

}
