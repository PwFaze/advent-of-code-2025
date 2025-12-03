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
		n1Idx := 0
		for i := range line {
			if i == len(line)-1 {
				break
			}
			if line[i] > line[n1Idx] {
				n1Idx = i
			}
		}
		n2Idx := n1Idx + 1
		for i := n2Idx; i < len(line); i++ {
			if line[i] > line[n2Idx] {
				n2Idx = i
			}
		}
		n1 := int(line[n1Idx] - '0')
		n2 := int(line[n2Idx] - '0')
		res += n1*10 + n2

	}
	return res
}

func Part02(r io.Reader) any {
	lines := util.ReadLines(r)
	wantedLength := 12
	var result int64
	for _, line := range lines {
		stack := []rune{}
		lineLength := len(line)
		for i, n := range line {
			num := n - '0'
			for len(stack) > 0 && stack[len(stack)-1] < num && len(stack)-1+lineLength-i >= wantedLength {
				// pop in these condition
				// 1.when len > 0
				// 2.current number > current top of stack
				// 3.if poped then the length should still reach 12(wantedLength) \
				// -> current stack length + lineLength - current number index >= wanted length
				stack = stack[:len(stack)-1]
			}
			// append in this condition
			if len(stack) < wantedLength {
				stack = append(stack, num)
			}

		}
		var sb strings.Builder
		for _, n := range stack {
			sb.WriteRune(n + '0')
		}

		num, _ := strconv.ParseInt(sb.String(), 10, 64)
		result += num
	}
	return result
}
