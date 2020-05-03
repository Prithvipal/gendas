package series

import (
	"github.com/Prithvipal/godf/series/types"
)

// IntElement ...
type IntElement struct {
	data types.IntType
}

// Set ...
func (e IntElement) Set(data interface{}) {
	x := data.(int)
	e.data = types.IntType(x)
	//fmt.Println(types.IntType)
	//fmt.Println(types.IntType(x))
}
