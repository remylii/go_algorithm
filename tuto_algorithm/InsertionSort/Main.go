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
	length := nextInt()
	var arr = make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = nextInt()
	}
	fmt.Println(arr)

	fmt.Println(arr)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
