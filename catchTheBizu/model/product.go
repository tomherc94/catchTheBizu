package model

import "fmt"

//Product contains the information needes to analyze opportunities
type Product struct {
	Code              string   `json:"code"`
	CurrentPrice      float64  `json:"currentPrice"`
	CeilingPrice      float64  `json:"ceilingPrice"`
	PossibleValuation float64  `json:"possibleValuation"`
	Segment           string   `json:"segment"`
	Category          Category `json:"category"`
}

//NewProduct instantiates a new product
func (p *Product) NewProduct(code, segment string, currentPrice, ceilingPrice float64, category Category) {

	p.Code = code
	p.CurrentPrice = currentPrice
	p.CeilingPrice = ceilingPrice
	p.Segment = segment
	p.Category = category
	p.PossibleValuation = p.CalculatePossibleValuation(p.CurrentPrice, p.CeilingPrice)
}

//CalculatePossibleValuation return the possible valuation
func (p *Product) CalculatePossibleValuation(currentPrice, ceilingPrice float64) float64 {
	possibleValuation := (ceilingPrice / currentPrice) - 1
	return possibleValuation
}

//ToString prints the product
func (p *Product) ToString() {
	fmt.Printf("Code: %s", p.Code)
	fmt.Printf("Current Price: %0.2f\n", p.CurrentPrice)
	fmt.Printf("Ceiling Price: %0.2f\n", p.CeilingPrice)
	fmt.Printf("Possible Valuation: %0.2f\n", p.PossibleValuation)
	fmt.Printf("Segment: %s\n", p.Segment)
	fmt.Printf("Category: %s", p.Category)
}
