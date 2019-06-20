package comm

import (
	"fmt"
)

func PrintIntSlice(info string, nums []int) {
	fmt.Print(info, ":")
	for _, n := range nums {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
