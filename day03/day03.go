package day03

import (
	"io"
	"strconv"
	"strings"

	"github.com/PwFaze/advent-of-code-2025/util"
)

func Part01(r io.Reader) any {
	lines := util.ReadLines(r)
	res := 0
	for _, line := range lines {
		curr := 0
		l := 0
		r := 1
		for l <= len(line)-2 {
			joltCharge, _ := strconv.Atoi(string(line[l]) + string(line[r]))
			if joltCharge > curr {
				curr = joltCharge
			}
			r += 1
			if r >= len(line) {
				l += 1
				r = l + 1
			}
		}
		res += curr
	}
	return res
}

func Part02(r io.Reader) any {
	lines := util.ReadLines(r)
	orderLine := lines[0]
	orders := strings.Split(orderLine, ",")
	res := 0
	for _, order := range orders {
		o := strings.Split(order, "-")
		start, _ := strconv.Atoi(o[0])
		end, _ := strconv.Atoi(o[1])
		for i := start; i <= end; i++ {
			curr := strconv.Itoa(i)
			l := len(curr)
			for j := 1; j <= l/2; j++ {
				partialCurr := curr[:j]
				repeated := l / j
				if s := strings.Repeat(partialCurr, repeated); s == curr {
					res += i
					break
				}
			}
		}

	}
	return res
}
