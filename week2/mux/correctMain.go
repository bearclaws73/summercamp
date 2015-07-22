package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "DOG")
}

type CatHandler int

func main() {
	var cat CatHandler
	var dog DogHandler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
	})

	http.HandleFunc("/dog.jpeg", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "dog.jpeg")
	})

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		dogName := strings.Split(req.URL.Path, "/")[2]
		io.WriteString(res, `<!DOCTYPE html>
  <html>
    <body>
      Dog Name: <strong>`+dogName+`</strong><br>
      <img src="/dog.jpeg">
    </body>
  </html>`)
	})

	http.ListenAndServe(":9000", nil)
}
