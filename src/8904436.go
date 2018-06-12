package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

var digitMap = map[byte]int{
	'1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
	'6': 6, '7': 7, '8': 8, '9': 9, 'A': 10,
	'C': 11, 'D': 12, 'E': 13, 'F': 14, 'H': 15,
	'J': 16, 'K': 17, 'L': 18, 'M': 19, 'N': 20,
	'P': 21, 'R': 22, 'T': 23, 'V': 24, 'W': 25, 'X': 26,
	'B': 8, 'G': 11, 'I': 1, 'S': 5, 'U': 24, 'Y': 24, 'Z': 2,
}
var position = []int{2, 4, 5, 7, 8, 10, 11, 13, 26}

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func next() []byte {
	in.Scan()
	return in.Bytes()
}

func main() {
	in.Split(bufio.ScanWords)
	for t := nextInt(); t > 0; t-- {
		main2()
	}
}

func main2() {
	k := nextInt()
	var r, s int
	for i, c := range next() {
		r *= 27
		r += digitMap[c]
		s += position[i] * digitMap[c]
	}
	if s%27 != 0 {
		fmt.Println(k, "Invalid")
	} else {
		fmt.Println(k, r/27)
	}
}
