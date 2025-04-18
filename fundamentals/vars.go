package main

import "fmt"

func main() {
	a := "Vivek"
	fmt.Println(a)
	d := 3.14135
	var e int = int(d)
	fmt.Println(e)
	pointerFunctios()
}

func arithematicOperations(a, b int) {
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}

func iotaFunctions() {

	const (
		a = iota
		b
		c = 3 * iota
	)

	fmt.Println("const values:", a, b, c)
}

func pointerFunctios() {
	a := "foo"
	b := &a

	fmt.Println(a, *b)
}
