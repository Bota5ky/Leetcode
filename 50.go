package temp
//与offer16相同，改进了calc2函数用位运算代替
func myPow2(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	return calc(x, n)
}

func calc2(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	res := calc(x, n>>1)
	if n&1 == 1 {
		return res * res * x
	}
	return res * res
}
