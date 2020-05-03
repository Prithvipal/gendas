package main

import (
	"fmt"

	"github.com/Prithvipal/godf/series"
	//"github.com/Prithvipal/godf/series"
	//"github.com/Prithvipal/godf/series/types"
)

func main() {

	se := series.New("Num", []int{123, 2134, 54, 34234})
	// for _, ele := range se.Elements {
	// 	fmt.Println(ele)
	// }
	fmt.Println(se)
}
