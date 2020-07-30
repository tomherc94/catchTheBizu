package model

//Product contains the information needes to analyze opportunities
type Product struct {
	Code              string   `json:"code"`
	CurrentPrice      float64  `json:"currentPrice"`
	CeilingPrice      float64  `json:"ceilingPrice"`
	PossibleValuation float64  `json:"possibleValuation"`
	Segment           string   `json:"segment"`
	Category          Category `json:"objective"`
}

//NewProduct instantiates a new product
func (p *Product) NewProduct(code, segment string, currentPrice, ceilingPrice, possibleValuation float64, category Category) {

	p.Code = code
	p.CurrentPrice = currentPrice
	p.CeilingPrice = ceilingPrice
	p.PossibleValuation = possibleValuation
	p.Segment = segment
	p.Category = category
}
