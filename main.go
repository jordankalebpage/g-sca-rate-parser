package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Record struct {
	jobCode     string
	jobTitle    string
	rate        float64
	description string
}

func main() {
	file := readFile("2025_SCA_Rates.csv")
	defer file.Close()

	records := readCSV(file, '|', true)

	for _, record := range records {
		log.Println(record.jobCode, record.jobTitle, record.rate, record.description)
	}
}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func readCSV(file *os.File, delimiter rune, skipHeaders bool) []Record {
	reader := csv.NewReader(file)
	reader.Comma = delimiter

	var records []Record

	// Skip the header row
	if skipHeaders {
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		rate, err := strconv.ParseFloat(record[2], 64)

		if err != nil {
			log.Fatal(err)
		}

		recordCode := record[0]
		jobTitle := record[1]

		records = append(records, Record{
			jobCode:  recordCode,
			jobTitle: jobTitle,
			rate:     rate,
		})
	}

	return records
}
