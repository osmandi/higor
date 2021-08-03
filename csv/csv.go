package csv

// CSV struct to parse CSV to DataFrame
type CSV struct {
	Sep            rune
	NaNLayout      string
	LazyQuotes     bool
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

// NaNLayout to set custom NaN format
func NaNLayout(nanLayout string) CSVOptions {
	return func(c *CSV) {
		c.NaNLayout = nanLayout
	}
}

// DatetimeLayout Set layout to parse datetime columns
func DatetimeLayout(datetimeLayout string) CSVOptions {
	return func(c *CSV) {
		c.DatetimeLayout = datetimeLayout
	}
}
