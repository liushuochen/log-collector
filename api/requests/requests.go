// Package requests is used to send an HTTP request.
// This file defined Requests structure and related functions and methods.
package requests

import (
	"api-server/config"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	HTTP  = "http://"
	HTTPS = "https://"
)

// Requests structure contains HTTP request properties. Please use New function to create and init this struct.
// - Name:  Request name.
// - Host:  Server hostname or IP address.
// - Port:  Port number.
// - Route: HTTP request route.
// - Method: HTTP request method. For example GET, POST, etc.
// - isVerify: A bool variable indicates verify HTTP request. Default is true.
// - Params: HTTP request params. Usually used in GET request.
// - Headers: HTTP request headers.
//           Default Headers attribute contains following key-value pairs:
//           User-Agent: LogCollector/config.ServiceVersion
//           Accept: */*
//           Accept-Encoding: gzip, deflate, br
//           Connection: keep-alive
// - Mode: HTTP request mode. Currently, support HTTP and HTTPS. Default is HTTPS.
// - Timeout: HTTP request timeout. Default is 5 seconds.
type Requests struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Route    string `json:"route"`
	Method   string `json:"method"`
	isVerify bool
	Params   map[string]string `json:"params"`
	Headers  map[string]string `json:"headers"`
	Mode     string            `json:"mode"`
	Timeout  time.Duration     `json:"timeout"`
}

// New function will create a Requests object and init it.
func New(name, host, route, method string, port int) *Requests {
	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}
	req := &Requests{
		Name:     name,
		Host:     host,
		Port:     port,
		Route:    route,
		Method:   method,
		isVerify: true,
		Params:   make(map[string]string),
		Headers:  make(map[string]string),
		Mode:     HTTPS,
	}

	req.Headers["User-Agent"] = fmt.Sprintf("LogCollector/%s", config.ServiceVersion)
	req.Headers["Accept"] = "*/*"
	req.Headers["Accept-Encoding"] = "gzip, deflate, br"
	req.Headers["Connection"] = "keep-alive"

	req.Timeout = time.Second * 5
	return req
}

// OpenVerify method used to open the HTTP request verify.
func (r *Requests) OpenVerify() {
	r.isVerify = true
}

// CloseVerify method used to close the HTTP request verify.
func (r *Requests) CloseVerify() {
	r.isVerify = false
}

// ClearParams method used to clear the Params map in Requests pointer.
func (r *Requests) ClearParams() {
	r.Params = make(map[string]string)
}

// ClearHeaders method used to clear the Headers map in Requests pointer, expect default headers.
func (r *Requests) ClearHeaders() {
	r.Headers = make(map[string]string)

	r.Headers["User-Agent"] = fmt.Sprintf("LogCollector/%s", config.ServiceVersion)
	r.Headers["Accept"] = "*/*"
	r.Headers["Accept-Encoding"] = "gzip, deflate, br"
	r.Headers["Connection"] = "keep-alive"
}

// Combine each Requests properties to a real HTTP request URL.
// For example: https://127.0.0.1:1234/health?id=001&name=demo
func (r *Requests) combineURL() string {
	urlList := []string{r.Mode, r.Host, ":", strconv.Itoa(r.Port), r.Route}
	paramsList := make([]string, 0)
	for key, value := range r.Params {
		paramsList = append(paramsList, key+"="+value)
	}
	if len(paramsList) > 0 {
		urlList = append(urlList, "?", strings.Join(paramsList, "&"))
	}

	return strings.Join(urlList, "")
}

// AddParam method used to add a key-value pair to Params map.
func (r *Requests) AddParam(key, value string) {
	r.Params[key] = value
}

// AddParams method used to add key-value pairs from a map argument to Params map.
func (r *Requests) AddParams(params map[string]string) {
	for key, value := range params {
		r.AddParam(key, value)
	}
}

// AddHeader method used to add a key-value pair to Headers map.
func (r *Requests) AddHeader(key, value string) {
	r.Headers[key] = value
}

// AddHeaders method used to add key-value pairs from a map argument to Headers map.
func (r *Requests) AddHeaders(headers map[string]string) {
	for key, value := range headers {
		r.AddHeader(key, value)
	}
}

// Send method used to send HTTP request.
func (r *Requests) Send() (*Response, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !r.isVerify,
		},
	}

	request, err := http.NewRequest(r.Method, r.combineURL(), nil)
	if err != nil {
		return nil, err
	}

	// add headers
	for key, value := range r.Headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   r.Timeout,
	}
	defer client.CloseIdleConnections()

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return newResponse(string(body), resp.StatusCode), nil
}
