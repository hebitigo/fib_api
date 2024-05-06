package utils

func Fibbonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibbonacci(n-1) + Fibbonacci(n-2)
}
