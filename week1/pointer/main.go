package main

//swaps 2 variables
import "fmt"

//Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

func swap(xPtr, yPtr *int) {

	temp := *xPtr //*x,*y=*y,*x
	*xPtr = *yPtr
	*yPtr = temp

}

func main() {
	x := 1
	y := 2
	swap(&x, &y)

	fmt.Println("x=", x)
	fmt.Println("y=", y) // x is 0
}
