package csv

import "strings"

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

// NaNLayout to set custom NaN format
func NaNLayout(nanLayout string) CSVOptions {
	return func(c *CSV) {
		c.NaNLayout = nanLayout
	}
}

// LazyQuotes True if the CSV has lazy quotes
func LazyQuotes(lazyQuotes bool) CSVOptions {
	return func(c *CSV) {
		c.LazyQuotes = lazyQuotes
	}
}

func parseDatetime(layout string) string {
	layout = strings.ToLower(layout)
	layout = strings.Replace(layout, "yyyy", "2006", -1)
	layout = strings.Replace(layout, "mm", "01", -1)
	layout = strings.Replace(layout, "dd", "02", -1)

	return layout
}

// TODO: CSVCreateMock
// TODO: ReadCSV (Sep, nanLayout, datetimeLayout)
