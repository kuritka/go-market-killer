package data

// https://github.com/kuritka/CommodityQuotations
// https://github.com/kuritka/MarketLoader
import "time"

type Category uint8

type Period uint8

var Categories = struct {
	Currencies        Category
	Energies          Category
	Financials        Category
	GrainsAndOilseeds Category
	Indices           Category
	Meats             Category
	Metals            Category
	Softs             Category
}{
	Currencies:        1,
	Energies:          2,
	Financials:        3,
	GrainsAndOilseeds: 4,
	Indices:           5,
	Meats:             6,
	Metals:            7,
	Softs:             8,
}

var Periods = struct {
	Unknown Period
	Hourly  Period
	Daily   Period
	Weekly  Period
	Monthly Period
}{
	Unknown: 0,
	Hourly:  1,
	Daily:   2,
	Weekly:  4,
	Monthly: 8,
}

type Quote struct {
	Symbol    string
	Period    Period
	DateTime  time.Time
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Volume    int
	Interrest int
}

func NewQuote() Quote {
	return Quote{Period: Periods.Unknown}
}

type Inflation struct {
	Annual  float64
	Year    int
	Monthly []float64
}

type Cpi struct {
	Year    int
	Monthly []float64
}

func (a Cpi) Annual() float64 {
	var sum float64
	for _, v := range a.Monthly {
		sum += v
	}
	return sum / float64(len(a.Monthly))
}

type ContractSpecification struct {
	Name          string
	SizeUnit      string
	ContractSize  float64
	Category      Category
	Exchange      string
	Months        string
	PointValueUsd float64
	Symbol        string
}
