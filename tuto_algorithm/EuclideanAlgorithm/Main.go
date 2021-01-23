package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	a := NextInt()
	b := NextInt()

	fmt.Println(a, b)

	Calc(a, b)
}

// Calc func
func Calc(a int, b int) {
	c := a % b
	if c == 0 {
		fmt.Printf("最大公約数は %d\n", b)
		return
	}

	Calc(b, c)
}

// NextInt function
func NextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
