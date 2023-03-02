package factorial

func Factorial(n int) int {
	product := 1
	for i := 1; i <= n; i++ {
		product *= i
	}
	return product
}
