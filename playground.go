package main

import "fmt"

type Salutation string

func main() {

	var messages Salutation = "Salutation!"
	fmt.Println(messages)

	var message string = "Hello Go World!"

	//var greeting *string = &message
	var a int
	a = 1
	var b int
	b = 1
	res := calcsomething(a, &message)
	fmt.Println(message, "a:", a, "b:", b, "res:", res)

	//	sl slice := { 'a', 'b'}
}

func calcsomething(a int, b *string) (res int) {
	a = 55
	*b = "wijziging"
	fmt.Println(a)
	return a
}
