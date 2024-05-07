package utils

import "math/big"

func Fibbonacci(n int) *big.Int {
	if n <= 0 {
		return nil
	} else if n == 1 {
		return big.NewInt(1)
	}

	prev := big.NewInt(1)
	current := big.NewInt(1)
	for i := 2; i < n; i++ {
		current, prev = new(big.Int).Add(current, prev), current
	}
	return current
}
