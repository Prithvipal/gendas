package series

import (
	"fmt"

	"github.com/Prithvipal/godf/series/types"
)

// New ...
func New(data interface{}, name string) Series {
	se := Series{
		Name: name,
	}

	switch data.(type) {
	case []int:
		se.Dtype = types.Int
		values := data.([]int)
		se.elements = make([]Element, len(values))
		for i, v := range values {
			intEle := IntElement{}
			intEle.Set(v)
			se.elements[i] = intEle
		}
	}

	fmt.Println(se)
	return se
}

// func getElements(values []interface{}) {
// 	eles := make([]Element, len(values))
// 	for i, v := range values {
// 		eles[i] = v
// 	}
// }
