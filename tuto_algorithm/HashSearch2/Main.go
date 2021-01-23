package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var recursiveCount int

func main() {
	sc.Split(bufio.ScanWords)

	length := nextInt()
	resLength := length + (length / 2) + 1
	arr, res := make([]int, length), make([]int, resLength) // 0 で初期化

	// 受け取った値を配列に初期化
	for i := 0; i < length; i++ {
		arr[i] = nextInt()
	}
	fmt.Println(arr)

	// ハッシュ配列作成
	for _, value := range arr {
		n := getHashArrayIndex(value, res)
		res[n] = value
	}
	fmt.Println(res)

	// 検索
	target := nextInt()
	hashSearch(target, res)
}

// ハッシュ関数
func myHashFunc(n int) int {
	return n % 11
}

// 再帰関数
func getHashArrayIndex(n int, res []int) int {
	index := myHashFunc(n)
	if res[index] == 0 {
		recursiveCount = 0
		return index
	}

	recursiveCount++
	if recursiveCount > len(res) {
		panic("ハッシュ配列に空きががない気がする")
	}

	return getHashArrayIndex(index+1, res)
}

// 探索
func hashSearch(target int, res []int) {

	fmt.Println(" > Search. ")
	c := 1

	targetIndex := myHashFunc(target)
	length := len(res)
	for i := 0; i < length; i++ {
		if res[targetIndex] == target {
			answer(target, targetIndex)
			break
		} else if res[targetIndex] == 0 {
			answer2(target)
			break
		}
		targetIndex = myHashFunc(targetIndex + 1)

		c++
	}
	fmt.Println(" > search count", c)
}

func answer(target int, targetIndex int) {
	fmt.Println(target, "index is", targetIndex)
}

func answer2(target int) {
	fmt.Println(target, "is not found")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
