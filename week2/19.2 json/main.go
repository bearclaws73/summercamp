package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Stockdata struct {
	X float64
	Y float64
}

func main() {

	f, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var obj []Stockdata

	err = json.NewDecoder(f).Decode(&obj) //or Stockdata?
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)

}
