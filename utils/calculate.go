package utils

import "math/big"

func Fibbonacci(n int) []*big.Int {
	if n <= 0 {
		return nil
	} else if n == 1 {
		return []*big.Int{big.NewInt(1)}
	}

	fib := make([]*big.Int, n)
	fib[0] = big.NewInt(1)
	fib[1] = big.NewInt(1)
	for i := 2; i < n; i++ {
		fib[i] = new(big.Int).Add(fib[i-1], fib[i-2])
	}
	return fib
}
