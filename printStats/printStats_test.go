package printStats

import (
	"fmt"
	"testing"
)

func TestPrintStats(t *testing.T) {
	var expected string
	var actual string
	lc := 1
	wc := 2
	cc := 5
	fname := "test.txt"
	oneFlagStr := "%7d %s\n"
	twoFlagStr := "%7d %7d %s\n"
	threeFlagStr := "%7d %7d %7d %s\n"

	// Single flag tests
	expected = fmt.Sprintf(oneFlagStr, lc, fname)
	actual = PrintStats(lc, wc, cc, fname, true, false, false, 1)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

	expected = fmt.Sprintf(oneFlagStr, wc, fname)
	actual = PrintStats(lc, wc, cc, fname, false, true, false, 1)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

	expected = fmt.Sprintf(oneFlagStr, cc, fname)
	actual = PrintStats(lc, wc, cc, fname, false, false, true, 1)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

	// Two flags
	expected = fmt.Sprintf(twoFlagStr, lc, wc, fname)
	actual = PrintStats(lc, wc, cc, fname, true, true, false, 2)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

	expected = fmt.Sprintf(twoFlagStr, lc, cc, fname)
	actual = PrintStats(lc, wc, cc, fname, true, false, true, 2)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

	expected = fmt.Sprintf(twoFlagStr, wc, cc, fname)
	actual = PrintStats(lc, wc, cc, fname, false, true, true, 2)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

	// Three flags
	expected = fmt.Sprintf(threeFlagStr, lc, wc, cc, fname)
	actual = PrintStats(lc, wc, cc, fname, true, true, true, 3)

	if expected != actual {
		t.Errorf("Output mismatch.\nExpected: %s\nGot: %s", expected, actual)
	}

}
