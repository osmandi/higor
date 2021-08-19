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

func TestDatetimeLayout(t *testing.T) {
	datetimeLayout := "yyyy-MM-dd"
	csvResult := &CSV{}
	csvOptionInternal := DatetimeLayout(datetimeLayout)
	csvOptionInternal(csvResult)

	if csvResult.DatetimeLayout != datetimeLayout {
		t.Errorf("Expected: %v\nResult: %v", datetimeLayout, csvResult.DatetimeLayout)
	}

}
