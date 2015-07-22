package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type monumentInformation struct {
	columns map[string]int
}

type monument struct {
	stationID string
	elevation string
}

func (info *monumentInformation) setColumns(record []string) {
	for idx, column := range record {
		info.columns[column] = idx
	}
}

func (info *monumentInformation) parseState(record []string) (*monument, error) {
	column := info.columns["id"]
	id, err := strconv.Atoi(record[column])
	if err != nil {
		return nil, err
	}
	stationID := record[info.columns["station id"]]
	elevation := record[info.columns["elevation"]]
	return &monument{
		stationID: stationID,
		elevation: elevation,
	}, nil
}

func main() {
	// #1 open a file

	_, err := os.Open("cityChampaign.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	highest := 0
	monumentLookup := map[string]*monument{}

	info := &monumentInformation{
		columns: make(map[string]int),
	}

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
			stationID, err := info.parseId(record)
			if err != nil {
				log.Fatalln(err)
			}
			monumentLookup[monument.station_id] = station_id
		}
	}

	// monument-information AL
	if len(os.Args) < 2 {
		log.Fatalln("expected monument id")
	}
	abbreviation := os.Args[1]
	monument, ok := monumentLookup[abbreviation]
	if !ok {
		log.Fatalln("invalid monument id")
	}

	fmt.Println(`
<html>
    <head></head>
    <body>
      <table>
        <tr>
          <th>Abbreviation</th>
          <th>Name</th>
        </tr>`)

	fmt.Println(`
        <tr>
          <td>` + monument.abbreviation + `</td>
          <td>` + monument.elevation + `</td>
        </tr>
    `)

	fmt.Println(`
      </table>
    </body>
</html>
    `)
}
