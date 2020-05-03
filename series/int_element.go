package series

import (
	"fmt"

	"github.com/Prithvipal/godf/series/types"
)

// IntElement ...
type IntElement struct {
	data types.IntType
}

// Set ...
func (e *IntElement) Set(data interface{}) {
	x := data.(int)
	e.data = types.IntType(x)
}

// Len ...
func (e *IntElement) Len() int {
	return len(e.String())
}

// Val ...
func (e *IntElement) Val() string {
	return e.String()
}

func (e IntElement) String() string {
	return fmt.Sprint(int(e.data))
}
