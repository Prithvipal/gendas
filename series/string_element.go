package series

import (
	"fmt"

	"github.com/Prithvipal/godf/series/types"
)

// StringElement ...
type StringElement struct {
	data types.StringType
}

// Set ...
func (e *StringElement) Set(data interface{}) {
	x := data.(string)
	e.data = types.StringType(x)
}

// Len ...
func (e *StringElement) Len() int {
	return len(e.String())
}

// Val ...
func (e *StringElement) Val() string {
	return e.String()
}

func (e StringElement) String() string {
	return fmt.Sprint(string(e.data))
}
