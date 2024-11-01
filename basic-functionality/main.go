package main

import (
	"fmt"

	"rsc.io/quote/v3"
)

func getQuote() string {
	return quote.HelloV3()
}

func getNewerQuote() string {
	return quote.Concurrency()
}

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
}
