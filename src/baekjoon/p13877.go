package baekjoon

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var T int
	numbers := make([]string, T)
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var tmp, a string
		fmt.Scanf("%s %s", &tmp, &a)
		a = strings.TrimLeft(a, "0")
		if a == "" {
			a = "0"
		}
		numbers = append(numbers, a)
	}
	// fmt.Println(numbers)
	for i, value := range numbers {
		if strings.Contains(value, "8") || strings.Contains(value, "9") {
			a, _ := strconv.ParseInt(value, 16, 64)
			fmt.Printf("%d 0 %s %d\n", i+1, value, a)
		} else {
			a, _ := strconv.ParseInt(value, 8, 64)
			b, _ := strconv.ParseInt(value, 16, 64)
			fmt.Printf("%d %d %s %d\n", i+1, a, value, b)
		}
	}
}
