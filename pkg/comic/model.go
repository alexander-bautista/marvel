package comic

// Comic :  comic model
type Comic struct {
	Id     int     `json:"id"`
	Title  string  `json:"title,omitempty"`
	Isbn   string  `json:"isbn,omitempty"`
	Format string  `json:"format,omitempty"`
	Dates  []date  `json:"dates"`
	Prices []Price `json:"prices"`
	Qty    int     `json:"quantity"`
}

type date struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type Price struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

const (
	TaxOverPrintPrice = 0.1
)

func (comic *Comic) EstimatedTaxes() (tax float64) {

	for _, t := range comic.Prices {
		if t.Type == "printPrice" {
			tax += t.Price * TaxOverPrintPrice
		}
	}
	// Another price types sums 0 on taxes
	return tax
}

/*func (comic Comic) String() string {
	return fmt.Sprintf("Comic %s %s. Quantity %d", comic.Title, comic.Format, comic.Qty)
}*/
