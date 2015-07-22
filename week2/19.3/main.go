// package main
//
// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )
//
// type row struct {
// 	Date time.Time
// 	Open float64
// 	y float64
// }
//
// func main() {
//
// 	src, err := os.Open("table.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer src.Close()
//
// 	dat,err :=os.Create("table.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer src.Close()
//
// 	rows,err :=csv.NewReader(src).ReadAll()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(rows)
//
// 	records := make([]record, 0)
//
//
//
//
// 	//json does not have time
// 	//csv.NewReader
// 	//conv each row to struct
// 	// each struct to a list of struct
// 	// encode that list to a file using json encoder
// 	var obj table
//
// 	err = json.NewDecoder(f).Decode(&obj) //or Stockdata?
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println(table)
//
// }

package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
	// High, Low, Close
}

func main() {
	src, err := os.Open("table.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("table.json")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	rows, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	records := make([]Record, 0, len(rows))
	for _, row := range rows {
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	err = json.NewEncoder(dst).Encode(records)
	if err != nil {
		panic(err)
	}

}
