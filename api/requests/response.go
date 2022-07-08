// Package requests is used to send a HTTP request.
// This file defined Response structure which returned from a HTTP request and related functions and methods.
package requests

// Response structure contains HTTP response properties.
// - Body: Response body.
// - Code: Response status code.
type Response struct {
	Body string `json:"body"`
	Code int    `json:"code"`
}

func newResponse(body string, code int) *Response {
	return &Response {
		Body: body,
		Code: code,
	}
}
