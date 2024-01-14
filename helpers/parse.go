package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCounts(fname string) (int, int, int) {
	var lc, wc, cc int
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
		// count characters
		cc += len(s)
		// count lines
		lc++
	}
	return lc, wc, cc
}
