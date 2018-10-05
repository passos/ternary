package ternary

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func Str(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func Byte(condition bool, trueVal, falseVal byte) byte {
	if condition {
		return trueVal
	}
	return falseVal
}

func Int(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

func Uint(condition bool, trueVal, falseVal uint) uint {
	if condition {
		return trueVal
	}
	return falseVal
}

func Float32(condition bool, trueVal, falseVal float32) float32 {
	if condition {
		return trueVal
	}
	return falseVal
}

func Float64(condition bool, trueVal, falseVal float64) float64 {
	if condition {
		return trueVal
	}
	return falseVal
}
