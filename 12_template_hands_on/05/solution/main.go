package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := parser("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}

}

func parser(filePath string) []Record {
	csvFile, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully Opened CSV file")

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(csvLines))

	for i, row := range csvLines {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records
}
