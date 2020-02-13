package comic

import "testing"

func TestEstimatedTaxesValidPrintPricesShouldNotBeZero(t *testing.T)  {
	comic := Comic{
		Id: 1,
	}
	var expected float32;
	for i := 0; i<5 ;i++ {
	  comic.Prices = append(comic.Prices, Price{Type: "printPrice", Price: 10})
	  expected += 10
	}
	expected = expected * 0.1
	got := comic.EstimatedTaxes();

	if comic.EstimatedTaxes() != expected {
		t.Errorf("EstimatedTaxes failed, expected %f, got %f", expected, got)
	}
}

func TestEstimatedTaxesShouldReturnZeroIfNoPrices(t *testing.T)  {
	comic := Comic{
		Id: 1,
	}
	var expected float32 =  0;

	got := comic.EstimatedTaxes();

	if comic.EstimatedTaxes() != expected {
		t.Errorf("EstimatedTaxes failed, expected %f, got %f", expected, got)
	}
}

func TestEstimatedTaxesShouldReturnZeroIfNoPrintPrices(t *testing.T)  {
	comic := Comic{
		Id: 1,
	}
	var expected float32 = 0;

	for i := 0; i<5 ;i++ {
		comic.Prices = append(comic.Prices, Price{Type: "comicPrice", Price: 20})
	}

	got := comic.EstimatedTaxes();

	if comic.EstimatedTaxes() != expected {
		t.Errorf("EstimatedTaxes failed, expected %f, got %f", expected, got)
	}
}