package util

func ToValue[T any](t *T) T {
	if t == nil {
		var def_ T
		return def_
	}
	return *t
}

func IntToBool(number *int) bool {
	if ToValue(number) > 0 {
		return true
	}
	return false
}
