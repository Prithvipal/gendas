package main

import (
	"fmt"

	"github.com/Prithvipal/godf/series"
)

func main() {
	se := series.New([]float64{1.1, 2.2, 3.4, 4.5}, "A")
	// for _, ele := range se.Elements {
	// 	fmt.Println(ele)
	// }
	fmt.Println(se)
}
