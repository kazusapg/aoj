package main

import "fmt"

func main() {
	var S int
	fmt.Scan(&S)
	ss := S % 60
	mm := (S - ss) % 3600 / 60
	hh := S / 3600
	fmt.Printf("%d:%d:%d\n", hh, mm, ss)
}
