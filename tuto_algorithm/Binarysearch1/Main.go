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
	sc.Split(bufio.ScanWords)
	length := nextInt()
	var slice []int
	for i := 0; i < length; i++ {
		slice = append(slice, nextInt())
	}
	fmt.Println(slice)

	target := nextInt()
	fmt.Println(target)

	head := 0
	tail := len(slice)
	t := 1
	for t > 0 {
		if head > tail {
			ans2()
			return
		}
		center := (head + tail) / 2
		fmt.Println(center)

		if target == slice[center] {
			ans(slice[center])
			t = -1
			return
		}

		if slice[center] < target {
			head = center + 1
		} else {
			tail = center - 1
		}
	}
}

func ans(r int) {
	fmt.Println(strconv.Itoa(r) + " is target")
}

func ans2() {
	fmt.Println("Not found")
}
