package main

import "fmt"

// const milesToKilometers = 1.60934 //milesToKilometers
//
// func main() {
// 	var miles float64
//
// 	fmt.Print("Enter miles: ")
// 	fmt.Scanf("%v", &miles)
// 	var kilos = miles * milesToKilometers
// 	fmt.Println(miles, "miles in Kilometers are", kilos)

// adds all the numbers divisible by 3 or 5
// func main() {
// 	total := 0
// 	for i := 1; i <= 1000; i++ {
// 		if i%3 == 0 || i%5 == 0 {
// 			total = total + i
// 		}
// 	}
// 	fmt.Println(total)
// }

// FizzBuzz: puts
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
