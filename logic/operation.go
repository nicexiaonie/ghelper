package logic

func Ternary(b bool, t interface{}, f interface{}) interface{} {
	if b {
		return t
	} else {
		return f
	}
}
