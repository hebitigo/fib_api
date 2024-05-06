package validator

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/hebitigo/fib_api/error_utils"
)

var (
	Validator *validator.Validate
	trans     ut.Translator
)

func init() {
	//https://github.com/go-playground/validator/blob/master/_examples/translations/main.go
	//多言語に対応するためのメッセージのうち、日本語のメッセージの設定のみをtransで取得し、
	//タグに対してメッセージを登録するんだと思う
	ja := ja.New()
	uni := ut.New(ja, ja)
	//globalなtransを:=で上書きしてしまうと、変数のシャドーイングが起こってしまうため
	//registerValidateMessageでtransに対して登録したメッセージが消えてしまう
	//そのため、transはグローバル変数として定義しておき、
	//init関数内では再代入のみを行うようにする
	trans, _ = uni.GetTranslator("ja")
	Validator = validator.New()

	//https://github.com/go-playground/validator/blob/master/_examples/custom-validation/main.go
	Validator.RegisterValidation("is-not-negative", NegativeCheck)

	registerValidateMessage(Validator, "is-not-negative", error_utils.ErrContentIsNegative, trans)
	registerValidateMessage(Validator, "required", error_utils.ErrContentNotSpecified, trans)
}

func registerValidateMessage(v *validator.Validate, tag string, err error, trans ut.Translator) {
	v.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}:"+err.Error(), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T(tag, fe.Field())
		if err != nil {
			//特にtagに対してメッセージが登録されていない場合は、fe.Error()を返す
			return fe.(error).Error()
		}
		return t
	})
}

func ResponseCustomMessage(c *gin.Context, errs validator.ValidationErrors) {
	var res error_utils.ValidateErr
	for _, e := range errs {
		log.Println(e.Translate(trans))
		res += error_utils.ValidateErr(fmt.Sprintf("%s\n", e.Translate(trans)))
	}
	error_utils.WriteErrorResponse(c, http.StatusBadRequest, error(res))
}

func NegativeCheck(fl validator.FieldLevel) bool {
	return fl.Field().Int() >= 0
}
