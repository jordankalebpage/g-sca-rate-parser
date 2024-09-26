package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type JobCode string

type Record struct {
	jobTitle    string
	rate        float64
	description string
}

func main() {
	file := readFile("2025_SCA_Rates.csv")
	defer file.Close()

	records := map[JobCode]Record{}

	records = readRatesCSV(file, records, '|', true)

	file = readFile("2023_sca_rates_export_arrs.csv")
	defer file.Close()
	records = readDescriptionsCSV(file, records, ',', true)

	// for jobCode, record := range records {
	// 	log.Println(jobCode, record.jobTitle, record.rate, record.description)
	// }
}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func readRatesCSV(
	file *os.File,
	records map[JobCode]Record,
	delimiter rune,
	skipHeaders bool,
) map[JobCode]Record {
	reader := csv.NewReader(file)
	reader.Comma = delimiter

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

		jobCode := JobCode(record[0])
		jobTitle := record[1]

		records[jobCode] = Record{
			jobTitle: jobTitle,
			rate:     rate,
		}

	}

	return records
}

func readDescriptionsCSV(
	file *os.File,
	records map[JobCode]Record,
	delimiter rune,
	skipHeaders bool,
) map[JobCode]Record {
	reader := csv.NewReader(file)
	reader.Comma = delimiter

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

		jobCode := JobCode(record[0])

		// check if job code exists
		if foundRecord, ok := records[jobCode]; ok {
			foundRecord.description = record[2]
			records[jobCode] = foundRecord
		} else {
			log.Println("job code not found:", jobCode)
		}
	}

	return records
}
