package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var mem runtime.MemStats

func main() {
	n := NextInt()
	Eratosthenes(n)
}

// Eratosthenes func
func Eratosthenes(n int) {
	PrintMemory()
	var arr = make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = 1
	}
	arr[0], arr[1] = 0, 0

	limit := int(math.Sqrt(float64(n)))
	fmt.Printf("Target count %d, sqrt %d, array length %d\n", n, limit, len(arr))

	for i := 2; i*i <= n; i++ {
		if arr[i] == 0 {
			continue
		}

		k := i
		for k <= n/i {
			arr[i*k] = 0
			k++
		}
	}

	// output
	for key, value := range arr {
		if value == 1 {
			fmt.Printf("%d ", key)
		}
	}
	fmt.Print("\n")
	PrintMemory()
}

// NextInt func
func NextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// PrintMemory func
func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}
