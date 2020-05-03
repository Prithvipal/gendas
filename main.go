package main

import (
	"fmt"

	"github.com/Prithvipal/godf/series"
)

func main() {
	se := series.New([]int{123, 2134, 54, 34234}, "Num")
	// for _, ele := range se.Elements {
	// 	fmt.Println(ele)
	// }
	fmt.Println(se)
}
