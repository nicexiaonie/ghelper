package ghelper

type logic struct {
}

var Logic logic

func (logic) Ternary(b bool, t interface{}, f interface{}) interface{} {
	if b {
		return t
	} else {
		return f
	}
}
