package helpers

import "fmt"

func CreateStatsString(lc, wc, cc int, fname string, printLc, printWc, printCc bool, flagCount int) string {
	oneFlagString := "%7d %s\n"
	twoFlagString := "%7d %7d %s\n"
	threeFlagString := "%7d %7d %7d %s\n"
	var output string
	switch {
	case printLc:
		switch flagCount {
		case 1:
			output = fmt.Sprintf(oneFlagString, lc, fname)
		case 2:
			switch {
			case printWc:
				output = fmt.Sprintf(twoFlagString, lc, wc, fname)
			case printCc:
				output = fmt.Sprintf(twoFlagString, lc, cc, fname)
			}
		case 3:
			output = fmt.Sprintf(threeFlagString, lc, wc, cc, fname)
		}
	case printWc:
		switch flagCount {
		case 1:
			output = fmt.Sprintf(oneFlagString, wc, fname)
		default:
			output = fmt.Sprintf(twoFlagString, wc, cc, fname)
		}
	default:
		output = fmt.Sprintf(oneFlagString, cc, fname)
	}

	return output
}
