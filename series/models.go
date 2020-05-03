package series

// Series ...
type Series struct {
	Name     string
	Elements Elements
	Dtype    string
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
