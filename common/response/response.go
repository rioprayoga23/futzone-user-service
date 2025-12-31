package response

import (
	"net/http"
	"user-service/constants"
	errConstant "user-service/constants/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

type ParamHttpResponse struct {
	Code    int
	Err     error
	Data    interface{}
	Message *string
	Gin     *gin.Context
	Token   *string
}

func HttpResponse(param ParamHttpResponse) {
	token := ""
	if param.Token != nil {
		token = *param.Token
	}

	if param.Err == nil {
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Success,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
			Token:   token,
		})
		return
	}

	message := errConstant.ErrInternalServer.Error()
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		message = param.Err.Error()
	}

	param.Gin.JSON(param.Code, Response{
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
	})

	return
}
