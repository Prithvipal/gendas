package series

import (
	"fmt"

	"github.com/Prithvipal/godf/series/types"
)

// FloatElement ...
type FloatElement struct {
	data types.FloatType
}

// Set ...
func (e *FloatElement) Set(data interface{}) {
	x := data.(float64)
	e.data = types.FloatType(x)
}

// String ...
func (e FloatElement) String() string {
	return fmt.Sprint(float64(e.data))
}
