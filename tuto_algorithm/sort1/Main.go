package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	var arr [7]int
	sc.Split(bufio.ScanWords)
	for i := 0; i < 7; i++ {
		arr[i] = nextInt()
	}
	fmt.Println(arr)

	t := 1
	cnt := 0
	length := len(arr)
	for t > 0 {
		cnt++

		t = -1
		for i := 0; i < length-1; i++ {
			fmt.Println(i, cnt)
			n := 0
			if arr[i] < arr[i+1] {
				n = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = n
				t = 1
			}
		}
	}

	fmt.Println(arr)
}
