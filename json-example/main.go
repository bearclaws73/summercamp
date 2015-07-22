package main

//keys are always string
import (
	"encoding/json"
	"fmt"
)

type Anything interface{}

func main() {
	jsonData := `
  {"key":"value"}
  `
	var obj map[string]interface{} //[] float64 instead of map... would work with 100.2,3.0 etc for jsonData
	//above for key:value
	err := json.Unmarshal([]byte(jsonData), &obj)

	if err != nil {
		panic(err)
	}
	fmt.Println("%T", obj) // %T gives you type

}
