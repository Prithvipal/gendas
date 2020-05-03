package series

import (
	"strings"
	"unicode/utf8"
)

// Series ...
type Series struct {
	Name     string
	Elements Elements
	Dtype    string
}

// Len ...
func (s Series) Len() int {
	return len(s.Elements)
}

// Records  ...
func (s Series) Records() []string {
	width := s.getColumnWidth()
	records := make([]string, s.Len()+1)
	records[0] = addLeftPadding(s.Name, width)
	for i, r := range s.Elements {
		record := addLeftPadding(r.Val(), width)
		records[i+1] = record
	}
	return records
}

func (s Series) String() string {
	records := s.Records()
	return strings.Join(records, "\n")
}

func addLeftPadding(s string, nchar int) string {
	if utf8.RuneCountInString(s) < nchar {
		return strings.Repeat(" ", nchar-utf8.RuneCountInString(s)+1) + s
	}
	return " " + s
}

// SetValues ...
// func (s Series) SetValues(values []types.BasicType) {
// 	for i, val := range values {
// 		if s.Dtype == types.Int {
// 			s.elements[i].Set(val)
// 		}
// 	}
// }

// Elements ...
type Elements []Element

// Element ...
type Element interface {
	Set(interface{})
	Len() int
	Val() string
}

// Conf ...
type Conf struct {
	inferName bool
}

// ConfOption ...
type ConfOption struct {
}

// SetConfOption ...
type SetConfOption func(*ConfOption)
