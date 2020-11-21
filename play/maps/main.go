package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./oscar_age_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nameCount := map[string]int{}
	for _, record := range records {
		nameCount[record[3]]++
	}

	for name, count := range nameCount {
		if count > 1 {
			fmt.Println(name, count)
		}
	}
}
