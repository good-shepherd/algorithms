package main

import (
	"bufio"
	"os"
	"fmt"
)

var in *bufio.Scanner
var n, l int

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func main() {
	file, err := os.Open("src/algospot/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	in = bufio.NewScanner(file)
	defer file.Close()
	// in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	a := nextInt()
	for t := 0; t < a; t++ {
		solution()
	}
}

func solution() {
	arr := make([]float64, 0)
	var min float64 = 10000
	n = nextInt()
	l = nextInt()
	for i := 0; i < n; i++ {
		arr = append(arr, float64(nextInt()))
	}
	for i := l; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			var sum float64 = 0
			for k := j; k < i+j; k++ {
				sum += arr[k]
			}
			avg := sum / float64(i)
			// fmt.Println("avg:", avg)
			if min > avg {
				min = avg
			}
		}
	}
	fmt.Printf("%.11f\n", min)
}
