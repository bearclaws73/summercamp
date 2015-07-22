package main

import "fmt"

func f(t string) {
	fmt.Println(t)
}

func main() {

	i := 1
	for i <= 5 {
		go f("hello world")
		i = i + 1
	}

	// go f("hello world")
	// go f("hello world")
	// go f("hello world")
	// go f("hello world")
	// go f("hello world")
	fmt.Scanln()
}
