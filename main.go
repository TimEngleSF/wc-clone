package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"wc-clone/printStats"
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
		parseFile(fname, &tlc, &twc, &tcc, *printLc, *printWc, *printCc, flagCount)
	}
	// if multiple files print a totals count
	if multiInput {
		if flagCount > 0 {
			output := printStats.PrintStats(tlc, twc, tcc, "total", *printLc, *printWc, *printCc, flagCount)
			fmt.Printf("%s", output)
		} else {
			fmt.Printf(" %7d %7d %7d %s\n", tlc, twc, tcc, "total")
		}
	}
}

func parseFile(fname string, tlc, twc, tcc *int, printLc, printWc, printCc bool, flagCount int) {
	var lc, wc, cc int
	file, err := os.Open(fname)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() { // this will loop over each line
		s := scan.Text()

		// split by whitespace to create slice of words and get the length
		wc += len(strings.Fields(s))
		*twc += wc
		// count characters
		cc += len(s)
		*tcc += cc
		// count lines
		lc++
		*tlc += lc
	}
	if flagCount > 0 {
		output := printStats.PrintStats(lc, wc, cc, fname, printLc, printWc, printCc, flagCount)
		fmt.Printf("%s", output)
	} else {
		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)
	}

	file.Close()
}
