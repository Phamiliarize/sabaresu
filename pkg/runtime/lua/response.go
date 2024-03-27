package lua

import "net/http"

type Response struct {
	StatusCode int
	Payload    interface{}
	Body       map[string]interface{}
	Header     http.Header
}
