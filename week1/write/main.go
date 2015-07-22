package main

import (
	"log"
	"os"
)

func main() {
	// Args[0] my-cat
	// Args[1]
	// $ my-cat

	f, err := os.Create("hello.txt")

	if err != nil {
		log.Fatalln("my program broke", err.Error())

	}
	defer f.Close()

	str := "txt"
	bs := []byte(str)

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file:", err.Error())

	}

}

//rdr:=strings.NewReader("something"
//
// bs, err := ioutil.ReadAll(f)
// if err != nil {
// 	log.Fatalln("my program broke")
