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
	var arr [4]int
	var res [7]int
	for i := 0; i < length; i++ {
		arr[i] = nextInt()
	}
	target := nextInt()
	fmt.Println(arr)

	for _, v := range arr {
		index := v % 7
		res[index] = v
	}
	fmt.Println(res)

	resIndex := target % 7
	fmt.Println(strconv.Itoa(target) + " target is " + strconv.Itoa(resIndex))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
