package main

import (
	"fmt"

	"math/rand"

	"rsc.io/quote/v3"
)

func getQuote() string {
	return quote.HelloV3()
}

func getNewerQuote() string {
	return quote.Concurrency()
}

type Person struct {
	firstName string
	lastName  string
}

type Season int

const (
	Spring Season = iota
	Summer
	Fall
	Winter
)

func main() {
	fmt.Println("Hi!")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 15
	if i < 10 {
		fmt.Println("Less than 10")
	} else if i < 20 {
		fmt.Println("Less than 20")
	}

	fmt.Println(getQuote())
	fmt.Println(getNewerQuote())

	p := Person{"Bobbo", "Jignesh"}
	fmt.Println(p.firstName + " " + p.lastName)

	c := rand.Intn(3)
	fmt.Println(c)
	switch c {
	case int(Spring):
		fmt.Println("Yay")
	case int(Summer):
		fmt.Println("Hot")
	case int(Fall):
		fmt.Println("Rain")
	default:
		fmt.Println("Bad")
	}

	fmt.Println("end")
}
