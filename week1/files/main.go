package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")

	if err != nil {
		log.Fatalln("my program broke", err.Error())

	}
	defer f.Close()

	//rdr:=strings.NewReader("something"

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke")
	}

	str := string(bs)

	fmt.Println(str)
}
