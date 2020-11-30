package router

import (
	"net/http"

	"stone/transaction-challenge/domain"
	"stone/transaction-challenge/transaction/payload"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TransactionHandler struct {
	TUsecase domain.TransactionUsecase
}

func NewTransactionRouter(router *gin.Engine, uc domain.TransactionUsecase) {
	handler := &TransactionHandler{
		TUsecase: uc,
	}

	planet := router.Group("/v1/transactions")
	{
		planet.POST("", handler.Insert)
	}
}

func (h *TransactionHandler) Insert(c *gin.Context) {
	var payload payload.InsertTransactionPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Unexpected JSON format"})
		return
	}

	transaction := payload.TransationPayload.ToEntity()

	var result map[string]interface{}
	if result, err = h.TUsecase.Insert(payload.AuthType, &transaction); err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternal:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
