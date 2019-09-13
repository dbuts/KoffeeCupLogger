package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Koffee Member Name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("What's good, " + text+"?")
	 }
