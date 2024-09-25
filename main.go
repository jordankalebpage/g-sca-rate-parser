package main

import (
	"log"
	"os"
)

type Record struct {
	jobCode     string
	jobTitle    string
	rate        float64
	description string
}

func main() {
	file, err := os.Open("2025_SCA_Rates.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records := make([]Record, 0)
}
