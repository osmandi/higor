package csv

import "testing"

// CSVOptions tests
func TestSep(t *testing.T) {
	sep := ';'
	csvResult := &CSV{}
	csvOptionInternal := Sep(sep)
	csvOptionInternal(csvResult)

	if csvResult.Sep != sep {
		t.Errorf("Sep error. Expected: ';'. But result: %v", csvResult.Sep)
	}
}

func TestNaNLayout(t *testing.T) {
	nanLayout := "nan"
	csvResult := &CSV{}
	csvOptionInternal := NaNLayout(nanLayout)
	csvOptionInternal(csvResult)

	if csvResult.NaNLayout != nanLayout {
		t.Errorf("NaN error. Expected: %v - But result: %v", nanLayout, csvResult.NaNLayout)
	}

}

func TestLazyQuotes(t *testing.T) {
	lazyQuotes := false
	csvResult := &CSV{}
	csvOptionInternal := LazyQuotes(lazyQuotes)
	csvOptionInternal(csvResult)

	if csvResult.LazyQuotes != lazyQuotes {
		t.Errorf("LazyQuotes error. Expected: %v - But result: %v", lazyQuotes, csvResult.LazyQuotes)
	}
}

func TestParseDatetimeLayout(t *testing.T) {
	// Datetime
	datetimeLayoutExpected := "2006-01-02"
	datetimeLayoutInput := "yyyy-mm-dd"
	datetimeLayoutResult := parseDatetime(datetimeLayoutInput)

	if datetimeLayoutExpected != datetimeLayoutResult {
		t.Errorf("Expected: %v\nResult: %v", datetimeLayoutExpected, datetimeLayoutResult)
	}

}
