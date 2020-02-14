package comic_test

import (
	"testing"

	"github.com/alexander-bautista/marvel/pkg/comic"
)

type TestComicData struct {
	comic    comic.Comic
	expected float32
}

func TestEstimatedTaxesShouldNotBeZeroIfValidPrintPrices(t *testing.T) {

	dataComics := []TestComicData{
		{
			comic: comic.Comic{
				Id: 1,
				Prices: []comic.Price{
					{
						Type: "printPrice", Price: 10,
					},
					{
						Type: "printPrice", Price: 20,
					},
				},
			},
			expected: 30 * comic.TaxOverPrintPrice,
		},
		{
			comic: comic.Comic{
				Id: 2,
				Prices: []comic.Price{
					{
						Type: "printPrice", Price: 15,
					},
					{
						Type: "printPrice", Price: 8,
					},
				},
			},
			expected: 23 * comic.TaxOverPrintPrice,
		},
	}

	for _, item := range dataComics {
		got := item.comic.EstimatedTaxes()

		if item.expected != got {
			t.Errorf("EstimatedTaxes FAILED, expected %f, got %f", item.expected, got)
		} else {
			t.Logf("EstimatedTaxes PASSED, expected %f, got %f", item.expected, got)
		}
	}
}

func TestEstimatedTaxesShouldReturnZeroIfNoPrices(t *testing.T) {
	comic := comic.Comic{
		Id: 1,
	}
	var expected float32 = 0

	got := comic.EstimatedTaxes()

	if comic.EstimatedTaxes() != expected {
		t.Errorf("EstimatedTaxes failed, expected %f, got %f", expected, got)
	}
}

func TestEstimatedTaxesShouldReturnZeroIfNoPrintPrices(t *testing.T) {
	item := comic.Comic{
		Id: 1,
	}
	var expected float32 = 0

	for i := 0; i < 5; i++ {
		item.Prices = append(item.Prices, comic.Price{Type: "comicPrice", Price: 20})
	}

	got := item.EstimatedTaxes()

	if item.EstimatedTaxes() != expected {
		t.Errorf("EstimatedTaxes failed, expected %f, got %f", expected, got)
	}
}
