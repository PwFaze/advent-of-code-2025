package util

import (
	"bufio"
	"io"
)

func ReadLines(r io.Reader) []string {
	lines := []string{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}

	return lines

}
