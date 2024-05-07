package handler

import (
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"

	"github.com/hebitigo/fib_api/error_utils"
	"github.com/hebitigo/fib_api/utils"
	val "github.com/hebitigo/fib_api/validator"
)

type FibonacciRequest struct {
	Count int `query:"n" validate:"required,is-not-negative"`
}

type FibonacciResponse struct {
	Result *big.Int `json:"result"`
}

func FibonacciHandler(c *gin.Context) {
	var req FibonacciRequest
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		log.Println("cannot convert query parameter to integer: ", err)
		error_utils.WriteErrorResponse(c, http.StatusBadRequest, error_utils.ErrContentCannotConvert)
		return
	}
	req.Count = n
	err = val.Validator.Struct(req)
	if err != nil {
		log.Println("validation error: ", err)
		if errs, ok := err.(validator.ValidationErrors); ok {
			val.ResponseCustomMessage(c, errs)
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	result, err := calculateFibonacciWithTimeout(ctx, req.Count)
	if err != nil {
		log.Println("error while calculating fibonacci: ", err)
		error_utils.WriteErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	res := FibonacciResponse{
		Result: result,
	}
	c.JSON(http.StatusOK, res)
}

func calculateFibonacciWithTimeout(ctx context.Context, count int) (*big.Int, error) {
	resultCh := make(chan *big.Int)
	go func() {
		result := utils.Fibbonacci(count)
		resultCh <- result
	}()
	select {
	case <-ctx.Done():
		return big.NewInt(0), error_utils.ErrConputationTooLong
	case result := <-resultCh:
		return result, nil
	}
}
