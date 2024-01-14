package parse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCounts(fname string, tlc, twc, tcc *int) (lc, wc, cc int) {
	// var lc, wc, cc int
	file, err := os.Open(fname)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer file.Close()

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
	return
}
