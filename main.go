package main

import (
	"fmt"

	"github.com/Prithvipal/godf/series"
)

func main() {
	se := series.New([]bool{false, true, true, false}, "A")
	// for _, ele := range se.Elements {
	// 	fmt.Println(ele)
	// }
	fmt.Println(se)
}
