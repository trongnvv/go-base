package gin_custom

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-base/pkg/helpers/log"
	"net/http"
)

type BaseController struct {
	validate *validator.Validate
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func newResponse(code int, message string, body interface{}) *response {
	return &response{
		Code:    code,
		Message: message,
		Body:    body,
	}
}

func NewBaseController(validate *validator.Validate) *BaseController {
	return &BaseController{validate}
}

func (b *BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Body:    data,
	})
}

func (b *BaseController) DefaultError(c *gin.Context, httpCode int) {
	c.JSON(httpCode, &response{
		Code:    httpCode,
		Message: http.StatusText(httpCode),
	})
}

func (b *BaseController) Error(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, &response{
		Code:    httpCode,
		Message: message,
	})
}

func (b *BaseController) ValidateRequest(request interface{}) error {
	err := b.validate.Struct(request)
	if err != nil {
		for _, errValidate := range err.(validator.ValidationErrors) {
			log.Debugf("query invalid, err:[%s]", errValidate)
		}
		return err
	}
	return nil
}
