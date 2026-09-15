package mathutils

func Factorial(n int) int64 {
	if n < 0 {
		return 0
	}
	var result int64 = 1
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

func Power(base int, exp int) int64 {
	var result int64 = 1
	b := int64(base)
	for i := 0; i < exp; i++ {
		result *= b
	}
	return result
}
