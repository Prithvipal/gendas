package types

// // BasicType ...
// type BasicType interface {
// 	ToString() string
// }

// IntType ...
type IntType int

// ToString ...
// func (t IntType) ToString() string {
// 	return ""
// }

type floatType float32

type stringType string

type boolType bool

const (
	// Int ...
	Int = "Int"
	//Float ...
	Float = "Float"
	// String ...
	String = "String"
	// Bool ...
	Bool = "Bool"
)
