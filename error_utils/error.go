package error_utils

import (
	"github.com/gin-gonic/gin"
)

type APIErr string

func (a APIErr) Error() string {
	return string(a)
}

type ValidateErr string

func (v ValidateErr) Error() string {
	return string(v)
}

const (
	ErrContentNotSpecified  ValidateErr = "必須のパラメータが渡されていません。もしくは数値以外が渡されています"
	ErrContentIsNegative    ValidateErr = "負の値は許可されていません"
	ErrContentCannotConvert APIErr      = "渡されたパラメータは上手く変換できませんでした。構造を確認してください"
	ErrConputationTooLong   APIErr      = "計算時間が長すぎます"
	ErrCannotBindQueryParam APIErr      = "クエリパラメータのデータ形式が不正です"
)

type ErrResopnse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteErrorResponse(c *gin.Context, status int, err error) {
	response := ErrResopnse{
		Status:  status,
		Message: err.Error(),
	}
	c.JSON(status, response)
}
