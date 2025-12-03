package day01

import (
	"io"
	"strconv"

	"github.com/PwFaze/advent-of-code-2025/util"
)

func Part01(r io.Reader) any {
	lines := util.ReadLines(r)
	d := 50
	res := 0
	for _, line := range lines {
		dir := line[:1]
		dist, _ := strconv.Atoi(line[1:])
		if dir == "R" {
			d = (d + dist) % 100
		} else {
			d = (d - dist + 100) % 100

		}
		if d == 0 {
			res += 1
		}
	}
	return res
}

func Part02(r io.Reader) any {
	lines := util.ReadLines(r)
	d := 50
	res := 0
	for _, line := range lines {
		dir := line[:1]
		dist, _ := strconv.Atoi(line[1:])
		res += dist / 100
		dist = dist % 100
		for range dist {
			if dir == "R" {
				d++
			} else {
				d--
			}
			if d > 99 {
				d = 0
			}
			if d < 0 {
				d = 99
			}
			if d == 0 {
				res += 1
			}
		}
	}
	return res
}
