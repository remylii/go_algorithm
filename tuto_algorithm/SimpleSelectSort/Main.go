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

	var indexMin int
	for i := 0; i < length-1; i++ {
		indexMin = i

		for j := i + 1; j < length; j++ {
			if arr[j] < arr[indexMin] {
				indexMin = j
			}
		}
		w := arr[i]
		arr[i] = arr[indexMin]
		arr[indexMin] = w

		fmt.Println(i, indexMin, arr)
	}
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
