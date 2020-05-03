package main

import (
	"fmt"

	"github.com/Prithvipal/godf/series"
)

func main() {
	se := series.New([]string{"A1", "B2", "C4", "D5"}, "A")
	// for _, ele := range se.Elements {
	// 	fmt.Println(ele)
	// }
	fmt.Println(se)
}
