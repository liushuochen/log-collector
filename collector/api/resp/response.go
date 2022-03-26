// Package resp define HTTP response constant and functions.
package resp

import "github.com/gin-gonic/gin"

// HTTP response code constant
const (
	Ok               = 200
	Created          = 201
	Accepted         = 202
	BadRequest       = 400
	AuthFailed       = 401
	Forbidden        = 403
	ResourceNotFound = 404
	InternalError    = 500
)

// HTTPResponse structure includes response information.
// - Code: HTTP response code
// - Success: Indicates whether the request was successful
// - Message: Response body. Usually a JSON or a string
type HTTPResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

// NewResponse function used to create a HTTPResponse pointer.
// It determines whether the request was successful based on the status value.
// If the value of status is less than 300, the function will set HTTPResponse.Success to true. Default is false.
func NewResponse(status int, message interface{}) *HTTPResponse {
	success := false
	if status < 300 {
		success = true
	}

	return &HTTPResponse{
		Code:    status,
		Success: success,
		Message: message,
	}
}

// SendResponse function used to send a HTTP response.
func SendResponse(c *gin.Context, status int, message interface{}) {
	c.JSON(status, NewResponse(status, message))
}
