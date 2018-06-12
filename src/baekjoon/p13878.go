package baekjoon

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
)

func main() {
	//start := time.Now()
	// s := [27]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "C", "D", "E", "F", "H", "J", "K", "L", "M", "N", "P", "R", "T", "V", "W", "X"}
	m := map[rune]float64{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 8,
		'C': 11,
		'D': 12,
		'E': 13,
		'F': 14,
		'G': 11,
		'H': 15,
		'I': 1,
		'J': 16,
		'K': 17,
		'L': 18,
		'M': 19,
		'N': 20,
		'O': 0,
		'P': 21,
		'Q': 0,
		'R': 22,
		'S': 5,
		'T': 23,
		'U': 24,
		'V': 24,
		'W': 25,
		'X': 26,
		'Y': 24,
		'Z': 2,
	}
	var T int
	fmt.Scanln(&T)
	UCN := make([]string, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		// var tmp, a string
		ss := strings.Fields(sc.Text())
		// fmt.Println(ss)
		// fmt.Scanf("%s %s", &tmp, &a)
		UCN = append(UCN, string(ss[1]))
	}

	for i, value := range UCN {
		//fmt.Println("value: " + value)
		a := int64(2*m[rune(value[0])]+4*m[rune(value[1])]+5*m[rune(value[2])]+7*m[rune(value[3])]+8*m[rune(value[4])]+10*m[rune(value[5])]+11*m[rune(value[6])]+13*m[rune(value[7])]) % 27
		last := int64(m[rune(value[len(value)-1])])
		//fmt.Printf("%d, %d\n", a, last)
		var sum float64
		if a != last {
			fmt.Printf("%d Invalid\n", i+1)
			continue
		} else {
			for j := 0; j < len(value)-1; j++ {
				sum += m[rune(value[len(value)-2-j])] * math.Pow(27, float64(j))
				//fmt.Println(int64(sum))
			}
			fmt.Printf("%d %d\n", i+1, int64(sum))
		}

	}
	//t := time.Now()
	//elapsed := t.Sub(start)
	//fmt.Println(elapsed)
}
