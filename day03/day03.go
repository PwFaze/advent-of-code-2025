package day03

import (
	"io"
	"strconv"
	"strings"

	"github.com/PwFaze/advent-of-code-2025/util"
)

func Part01(r io.Reader) any {
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
			if len(curr)%2 != 0 {
				continue
			}

			l := len(curr) / 2
			if curr[:l] == curr[l:] {
				res += i
			}
		}

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
