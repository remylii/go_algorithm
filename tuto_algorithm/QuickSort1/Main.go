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
	length := NextInt()
	var arr = make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = NextInt()
	}
	fmt.Println(arr)

	QuickSort(arr, 0, length-1)
	fmt.Println(arr)
}

// NextInt 入力受け取り
func NextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// QuickSort 実行アルゴリズム
func QuickSort(arr []int, left int, right int) {
	fmt.Println(">>> Left, Right", left, right)
	i, k := left+1, right
	for i < k {
		for 1 > 0 {
			if arr[i] < arr[left] && i < right {
				i++
				continue
			}
			fmt.Println("基準値より大きい値をみつけた i", i, arr[i])
			break
		}

		for 1 > 0 {
			if arr[k] >= arr[left] && k > left {
				k--
				continue
			}
			fmt.Println("基準値より小さい値をみつけた k", k, arr[k])
			break
		}

		// change
		if i < k {
			w := arr[k]
			arr[k] = arr[i]
			arr[i] = w
		}
	}
	fmt.Println("i, k, arr", i, k, arr)

	// 基準値を k 番目と入れ替える
	if arr[left] > arr[k] {
		w := arr[left]
		arr[left] = arr[k]
		arr[k] = w
	}

	if left < k-1 {
		QuickSort(arr, left, k-1)
	}
	if k+1 < right {
		QuickSort(arr, k+1, right)
	}
	// return arr
}
