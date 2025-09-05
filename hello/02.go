package hello

import (
	"fmt"
)

func main() {
	var n, result int
	fmt.Scan(&n)
	for n > 0 {
		n = n >> 1
		result++
	}
	fmt.Print(result)
}
