package error_utils

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIErr_Error(t *testing.T) {
	tests := []struct {
		name string
		a    APIErr
		want string
	}{
		{name: "Test APIErr",
			a:    ErrConputationTooLong,
			want: "計算時間が長すぎます",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.Error())
		})
	}
}

func TestValidateErr_Error(t *testing.T) {

	tests := []struct {
		name string
		v    ValidateErr
		want string
	}{{
		name: "Test ValidateErr Error",
		v:    ErrContentNotSpecified,
		want: "必須のパラメータが渡されていません。もしくは数値以外が渡されています",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.v.Error())
		})
	}
}

func TestWriteErrorResponse(t *testing.T) {
	type args struct {
		c      *gin.Context
		status int
		err    error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteErrorResponse(tt.args.c, tt.args.status, tt.args.err)
		})
	}
}
