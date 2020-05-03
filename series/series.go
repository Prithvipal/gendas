package series

import (
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
		se.Elements = make([]Element, len(values))
		for i, v := range values {
			intEle := &IntElement{}
			intEle.Set(v)
			se.Elements[i] = intEle
		}
	case []float64:
		se.Dtype = types.Float
		values := data.([]float64)
		se.Elements = make([]Element, len(values))
		for i, v := range values {
			floatEle := &FloatElement{}
			floatEle.Set(v)
			se.Elements[i] = floatEle
		}
	case []string:
		se.Dtype = types.String
		values := data.([]string)
		se.Elements = make([]Element, len(values))
		for i, v := range values {
			stringEle := &StringElement{}
			stringEle.Set(v)
			se.Elements[i] = stringEle
		}
	}

	//fmt.Println(se)
	return se
}

// func getElements(values []interface{}) {
// 	eles := make([]Element, len(values))
// 	for i, v := range values {
// 		eles[i] = v
// 	}
// }
