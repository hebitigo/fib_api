package error_utils

import (
	"fmt"
	"net/http/httptest"
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
	}{
		{
			name: "Test ValidateErr Error",
			v:    ErrContentNotSpecified,
			want: "必須のパラメータが渡されていません",
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
		status int
		err    error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test WriteErrorResponse",
			args: args{
				status: 400,
				err:    ErrContentCannotConvert,
			},
			want: fmt.Sprintf(`{"status":%d,"message":"%s"}`, 400, string(ErrContentCannotConvert)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			WriteErrorResponse(c, tt.args.status, tt.args.err)
			assert.Equal(t, tt.want, rec.Body.String())
		})
	}
}
