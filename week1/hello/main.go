package main

// added .2 before f in %f, has error

import (
	"fmt"
	"os"
	"strconv"
)



	fmt.Println("<!DocType html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")
	var kilos = number * miToKilo
	fmt.Println("+--------------------")
	fmt.Printf("Miles:%.2f<br>\n", number)

	fmt.Println("            |")
	fmt.Println("+--------------------")
	fmt.Printf("Kilometers:%.2f<br>\n", kilos)
	fmt.Println("    |")

	fmt.Println("+--------------------")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

//
// main(){  //different kind of if
// if value,ok:=somefunction(); !ok{
// 		panic("this should not happen")
//    }
// )

//
// func main() {  //formatting practice for float64
// 	fmt.Printf("%.2f", 1.23)
//
// 	fmt.Sprint("%.4f", 1.23)
// }

//milesToKilometers
// const milesToKilometers = 1.60934 //milesToKilometers
//
// func main() {
// 	var miles float64
//
// 	fmt.Print("Enter miles: ")
// 	fmt.Scanf("%v", &miles)
// 	var kilos = miles * milesToKilometers
// 	fmt.Println(miles, "miles in Kilometers are", kilos)
