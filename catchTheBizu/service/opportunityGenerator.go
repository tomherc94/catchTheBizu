package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/tomherc94/catchTheBizu/catchTheBizu/model"
)

func generateProducts(csvPath string) []model.Product {
	p := new(model.Product)
	products := []model.Product{}

	// Load a csv file.
	file, _ := os.Open(csvPath)

	// Create a new reader.
	reader := csv.NewReader(file)

	defer file.Close()

	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))
		for _, value := range record {
			attributes := strings.Split(value, ",")

			p.Code = attributes[0]
			p.CurrentPrice, err = strconv.ParseFloat(attributes[1], 64)
			p.CeilingPrice, err = strconv.ParseFloat(attributes[2], 64)
			p.CalculatePossibleValuation(p.CurrentPrice, p.CeilingPrice)

			if err != nil {
				panic(err)
			}

			products = append(products, *p)
		}
	}

	return products
}

func generateOpportunity([]model.Product) []model.Product {

	opportunities := []model.Product{}

	//ordena slice por valorização, categoria

	return opportunities
}
