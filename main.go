package main

import (
	"fmt"

	"github.com/Prithvipal/godf/series"
)

func main() {
	se := series.New([]int{1, 2, 3, 4}, "A")
	fmt.Println(se)
}
