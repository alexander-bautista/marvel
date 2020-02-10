package comic

import "fmt"

// Comic :  comic model
type Comic struct {
	Id     int     `json:"id"`
	Title  string  `json:"title,omitempty"`
	Isbn   string  `json:"isbn,omitempty"`
	Format string  `json:"format,omitempty"`
	Dates  []date  `json:"dates"`
	Prices []price `json:"prices"`
	Qty    int     `json:"quantity"`
}

type date struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type price struct {
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}

func (comic *Comic) EstimatedTaxes() (tax float32) {

	for _, t := range comic.Prices {
		if t.Type == "printPrice" {
			tax += t.Price * 0.1
		}
	}
	// Another price types sums 0 on taxes
	return tax
}

func (comic Comic) String() string {
	return fmt.Sprintf("Comic %s %s. Quantity %d", comic.Title, comic.Format, comic.Qty)
}
