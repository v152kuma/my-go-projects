package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	commandLineInputFunction()

}

func commandLineInputFunction() {

	fmt.Println("Please give me something to scream?")

	in := bufio.NewReader(os.Stdin)

	x, _ := in.ReadString('\n')

	x = strings.TrimSpace(x)

	x = strings.ToUpper(x)

	fmt.Println(x + "!")

}
