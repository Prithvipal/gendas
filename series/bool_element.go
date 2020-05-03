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

// Len ...
func (e *BoolElement) Len() int {
	return len(e.String())
}

// Val ...
func (e *BoolElement) Val() string {
	return e.String()
}

func (e BoolElement) String() string {
	return fmt.Sprint(bool(e.data))
}
