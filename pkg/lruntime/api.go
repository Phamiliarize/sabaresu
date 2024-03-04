package lruntime

import (
	"net/http"
	"net/url"

	lua "github.com/yuin/gopher-lua"
)

type Request struct {
	RequestID  string
	Method     string
	URL        *url.URL
	Header     http.Header
	Host       string
	gPathParam func(name string) string
}

type Response struct {
	StatusCode int
	Payload    interface{}
	Body       interface{}
	Header     http.Header
}

/* Functions to bind/expose */

func (r Request) getPathParam(L *lua.LState) int {
	lv := L.ToString(1)
	value := lua.LString(r.gPathParam(lv))
	L.Push(value)
	return 1
}

// MakeLRequest binds the native GoLang sabaresu request object into the Lua state
func MakeLRequest(L *lua.LState, req Request) *lua.LTable {
	lrequest := L.NewTable()
	lrequest.RawSetString("requestId", lua.LString(req.RequestID))
	lrequest.RawSetString("method", lua.LString(req.Method))
	lrequest.RawSetString("url", lua.LString(req.URL.String()))
	lrequest.RawSetString("path", lua.LString(req.URL.Path))
	lrequest.RawSetString("host", lua.LString(req.Host))
	lrequest.RawSetString("getPathParam", L.NewFunction(req.getPathParam))

	return lrequest
}

// MakeLResponse binds the native GoLang sabaresu request object into the Lua state
func MakeLResponse(L *lua.LState, resp Response) *lua.LTable {
	lresponse := L.NewTable()
	lresponse.RawSetString("statusCode", lua.LNumber(200))

	return lresponse
}
