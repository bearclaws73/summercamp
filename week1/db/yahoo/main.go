package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type yahooInformation struct {
	columns map[string]int
}

type yahoot struct {
	date     string
	open     string
	high     string
	low      string
	closev   string
	volume   string
	adjClose string
}

func (info *yahooInformation) setColumns(record []string) {
	for idx, column := range record {
		info.columns[column] = idx
	}
}

func (info *yahooInformation) parseyahoot(record []string) (*yahoot, error) {
	date := record[info.columns["Date"]]
	open := record[info.columns["Open"]]
	high := record[info.columns["High"]]
	low := record[info.columns["Low"]]
	closev := record[info.columns["Close"]]
	volume := record[info.columns["Volume"]]
	adjClose := record[info.columns["Adj Close"]]
	//	low := rec
	//ord[info.columns["census_region_name"]]
	return &yahoot{
		date:     date,
		open:     open,
		high:     high,
		low:      low,
		closev:   closev,
		volume:   volume,
		adjClose: adjClose,

		//	low:  low,
	}, nil
}

func main() {
	// #1 open a file
	f, err := os.Open("yahootable.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	info := &yahooInformation{
		columns: make(map[string]int),
	}

	fmt.Println(`
  <html>
      <head></head>
      <body>
        <table>
          <tr>
            <th>Date</th>
            <th>Open</th>
            <th>High</th>
            <th>Low</th>
            <th>Close</th>
            <th>Volume</th>
            <th>Adj Close</th>

          </tr>`)
	// #2 parse a csv file
	csvReader := csv.NewReader(f)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			info.setColumns(record)
		} else {
			yahoot, err := info.parseyahoot(record)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(`
          <tr>
            <td>` + yahoot.date + `</td>
            <td>` + yahoot.open + `</td>
            <td>` + yahoot.high + `</td>
            <td>` + yahoot.low + `</td>
            <td>` + yahoot.closev + `</td>
            <td>` + yahoot.volume + `</td>
            <td>` + yahoot.adjClose + `</td>
          </tr>
      `)
		}
	}

	fmt.Println(`
      </table>
    </body>
</html>
    `)
}
