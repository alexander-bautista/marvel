package comic_test

import (
	"math"
	"testing"

	"github.com/alexander-bautista/marvel/pkg/comic"
	"github.com/google/go-cmp/cmp"
)

type TestComicData struct {
	comic comic.Comic
	name  string
	want  float64
}

func TestEstimatedTaxes(t *testing.T) {

	const tolerance = .00001
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0

		if math.IsNaN(delta / mean) {
			return true
		}

		return delta/mean < tolerance
	})

	tests := []TestComicData{
		{
			name: "sucess with low prices",
			comic: comic.Comic{
				Id: 1,
				Prices: []comic.Price{
					{
						Type: "printPrice", Price: 10.45,
					},
					{
						Type: "printPrice", Price: 20.65,
					},
				},
			},
			want: (10.45 + 20.65) * comic.TaxOverPrintPrice,
		},
		{
			name: "sucess with large prices",
			comic: comic.Comic{
				Id: 2,
				Prices: []comic.Price{
					{
						Type: "printPrice", Price: 15456987,
					},
					{
						Type: "printPrice", Price: 5699982348,
					},
				},
			},
			want: (15400000 + 5699982348) * comic.TaxOverPrintPrice,
		},
		{
			name: "should return zero if no prices",
			comic: comic.Comic{
				Id: 2,
			},
			want: 0,
		},
		{
			name: "should return zero if no print prices",
			comic: comic.Comic{
				Id: 2,
				Prices: []comic.Price{
					{
						Type: "anotherPrice", Price: 15456987,
					},
					{
						Type: "somePrice", Price: 5699982348,
					},
				},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.comic.EstimatedTaxes(); !cmp.Equal(got, tt.want, opt) {
				t.Errorf("EstimatedTaxes FAILED, want %f, got %f", tt.want, got)
			} else {
				t.Logf("EstimatedTaxes PASSED, want %f, got %f", tt.want, got)
			}
		})
	}
}

/*
Approximate equality for floats can be handled by defining a custom comparer on floats that determines two values to be equal
if they are within some range of each other.
*/
