package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
import "log"

func MultiReader(readers...Reader) Reader {
	r := make([]Reader, len(readers))
	copy(r, readers)
	return &multiReader{r}

	for _, v := range os.Args[1:] {
	}
}
package main

import (
	"io"
	"os"
)
import "log"

func main() {
	// my-cat f1 f2 f3 ...
	// os.Args[0] == my-cat
	// os.Args[1] == f1
	// os.Args[2] == f2
	// os.Args[3] == f3
	rdrs := make([]io.Reader, len(os.Args)-1)
	for i, v := range os.Args[1:] {
		f, err := os.Open(v)
		if err != nil {
			log.Fatalln("my program broke", err.Error())
		}
		defer f.Close()
		rdrs[i] = f
	}
	rdr := io.MultiReader(rdrs...)

	// $ my-cat
	// Enter File Name:

	//rdr := strings.NewReader("some string")
	io.Copy(os.Stdout, rdr)
}
