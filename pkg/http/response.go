package http

import (
	"github.com/gin-gonic/gin"
)

type Errno struct {
	Code    int
	Message string
}

type BasicResponseJSON struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Object  interface{} `json:"result"`
}

type ResponseJSON struct {
	BasicResponseJSON
	TotalCount interface{} `json:"total_count"`
}

var (
	OK          = 200
	Error       = 400
	NotFound    = 404
	ServerError = 500
)

func NormalResponse(c *gin.Context, code int, message string, object interface{}, total int) {
	var resp ResponseJSON
	resp = ResponseJSON{BasicResponseJSON: BasicResponseJSON{Code: code, Message: message, Object: object}, TotalCount: total}
	c.JSON(200, resp)
}

func SimpleResponse(c *gin.Context, code int, message string, object interface{}) {
	var resp BasicResponseJSON
	resp = BasicResponseJSON{Code: code, Message: message, Object: object}
	c.JSON(200, resp)
}

func BadResponse(c *gin.Context, code int, message string) {
	var resp BasicResponseJSON
	resp = BasicResponseJSON{Code: code, Message: message, Object: nil}
	c.JSON(200, resp)
}
