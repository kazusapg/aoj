package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	a, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}
func scanSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func isInsideSquare(W, H, x, y, r int) bool {
	xStart := x - r
	xEnd := x + r
	if xStart < 0 || W < xEnd {
		return false
	}
	yStart := y - r
	yEnd := y + r
	if yStart < 0 || H < yEnd {
		return false
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	s := scanSlice(5)
	W, H, x, y, r := s[0], s[1], s[2], s[3], s[4]
	if isInsideSquare(W, H, x, y, r) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
