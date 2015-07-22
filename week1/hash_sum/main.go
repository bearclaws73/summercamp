package main

// This has code built in to print everything when the code is finished rather than by user input
//We use channels instead of changing variables simultaneously
have multiple steps and put channels between steps
import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func md5file(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	return h.Sum(nil)
}

func printmd5(fileName string) {   //
	bs := md5file(fileName)
	fmt.Printf("%s %x\n", fileName, bs)

	wg.Done()
}
func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go printmd5(path)
		return nil
	})
	wg.Wait()

}

/*
walk is recursive
readdir is not
*/
