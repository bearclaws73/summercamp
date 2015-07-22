package main

//reading from one file and writing to another file
import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// fromFileName :=os.Args[1]
	// toFileName :=os.Args[2]

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("my program broke", err.Error())

	}
	defer f.Close()

	//rdr:=strings.NewReader("something"

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke")
	}

	tf, err := os.Create(os.Args[2])

	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer tf.Close()

	_, err = tf.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file:", err.Error())

	}
}
