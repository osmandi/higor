package csv

// CSV struct to parse CSV to DataFrame
type CSV struct {
	Sep            rune
	LazyQuotes     bool
	NaNLayout      string
	DatetimeLayout string
}

// CSVOptions To apply optionals parameters
type CSVOptions func(csv *CSV)

// CSVOptions functions

// Sep to set custom separator
func Sep(sep rune) CSVOptions {
	return func(c *CSV) {
		c.Sep = sep
	}
}

func NaNLayout(nanLayout string) CSVOptions {
	return func(c *CSV) {
		c.NaNLayout = nanLayout
	}
}

// TODO: CSVCreateMock
// TODO: ReadCSV (Sep, nanLayout, datetimeLayout)
