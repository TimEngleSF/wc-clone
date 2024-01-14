package main

import (
	"flag"
	"fmt"
	"os"
	"wc-clone/helpers"
)

func main() {
	twc := 0
	tcc := 0
	tlc := 0
	multiInput := false

	printLc := flag.Bool("l", false, "print line count")
	printWc := flag.Bool("w", false, "print word count")
	printCc := flag.Bool("c", false, "print char count")
	flag.Parse()
	flagCount := flag.NFlag()

	if len(flag.Args()) > 1 {
		multiInput = true
	}
	// loop over file arguments and parse them to obtain their counts
	for _, fname := range os.Args[1:] {
		if fname == "-l" || fname == "-w" || fname == "-c" {
			continue
		}

		lc, wc, cc := helpers.GetCounts(fname)
		tlc += lc
		twc += wc
		tcc += cc

		if flagCount > 0 {
			output := helpers.CreateStatsString(lc, wc, cc, fname, *printLc, *printWc, *printCc, flagCount)
			fmt.Printf("%s", output)
		} else {
			fmt.Printf("%7d %7d %7d %s\n", tlc, twc, tcc, fname)
		}
	}
	// if multiple files print a totals count
	if multiInput {
		if flagCount > 0 {
			output := helpers.CreateStatsString(tlc, twc, tcc, "total", *printLc, *printWc, *printCc, flagCount)
			fmt.Printf("%s", output)
		} else {
			fmt.Printf("%7d %7d %7d %s\n", tlc, twc, tcc, "total")
		}
	}
}
