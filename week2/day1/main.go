//examples/week2/example-main.go
package main

import (
	"examples/week2/day1/example"
)

func main() {
	example.Hello()
}

// package main
//
// import (
// 	"fmt"
// 	"io"
// 	"net"
// 	"time"
// )
//
// func main() {
// 	ln, err := net.Listen("tcp", ":9000")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ln.Close()
//
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		io.WriteString(conn, fmt.Sprint(time.Now()))
//
// 		conn.Close()
// 	}
// }
//
//
