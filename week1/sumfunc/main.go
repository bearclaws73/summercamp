package main

import "fmt"

func first(x int) (int, bool) {
	if x%2 == 0 {
		return x / 2, true
	} else {
		return x / 2, false
	}
}

func main() {
	x := 71
	y, hx := first(x)
	fmt.Println(y, hx)

}

//return n/2, n%2==0
