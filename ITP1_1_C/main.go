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
func main() {
	sc.Split(bufio.ScanWords)
	y := scanSlice(2)
	a, b := y[0], y[1]
	area := a * b
	perimeter := (a * 2) + (b * 2)
	fmt.Println(area, perimeter)
}
