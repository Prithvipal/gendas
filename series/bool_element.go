package series

import (
	"fmt"

	"github.com/Prithvipal/godf/series/types"
)

// BoolElement ...
type BoolElement struct {
	data types.BoolType
}

// Set ...
func (e *BoolElement) Set(data interface{}) {
	x := data.(bool)
	e.data = types.BoolType(x)
}

func (e BoolElement) String() string {
	return fmt.Sprint(bool(e.data))
}
