package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1から始まるようにしていされて、かつ99の時の値が
// validationで不正な値は全て弾くので、正の値のみを受け付ける
func TestFibbonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test Fibbonacci 1",
			args: args{n: 1},
			want: big.NewInt(1),
		},
		{
			name: "Test Fibbonacci 25",
			args: args{n: 99},
			want: func() *big.Int{
				bI, _ := new(big.Int).SetString("218922995834555169026",10)
				return bI
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Fibbonacci(tt.args.n)[tt.args.n-1])
		})
	}
}
