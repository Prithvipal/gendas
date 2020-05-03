package series

import (
	"github.com/Prithvipal/godf/series/types"
)

// New ...
func New(name string, data interface{}) Series {
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
	case []bool:
		se.Dtype = types.Bool
		values := data.([]bool)
		se.Elements = make([]Element, len(values))
		for i, v := range values {
			boolEle := &BoolElement{}
			boolEle.Set(v)
			se.Elements[i] = boolEle
		}
	}

	//fmt.Println(se)
	return se
}

func (ser Series) getColumnWidth() int {
	columnWidth := len(ser.Name)
	for _, v := range ser.Elements {
		if v.Len() > columnWidth {
			columnWidth = v.Len()
		}
	}
	return columnWidth
}
