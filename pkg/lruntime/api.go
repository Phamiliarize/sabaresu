package lruntime

import (
	"net/http"
	"net/url"

	"github.com/Phamiliarize/sabaresu/pkg/lruntime/util"
	lua "github.com/yuin/gopher-lua"
)

type Request struct {
	ID         string
	Method     string
	URL        *url.URL
	Header     http.Header
	Host       string
	Body       map[string]interface{}
	gPathParam func(name string) string
}

type Response struct {
	StatusCode int
	Payload    interface{}
	Body       map[string]interface{}
	Header     http.Header
}

/* Functions to bind/expose */

func (r Request) getPathParam(L *lua.LState) int {
	lv := L.ToString(1)
	pathParam := r.gPathParam(lv)
	value := lua.LNil

	if pathParam != "" {
		value = lua.LString(r.gPathParam(lv))
	}

	L.Push(value)
	return 1
}

// MakeLRequest binds the native GoLang sabaresu request object into the Lua state
func MakeLRequest(L *lua.LState, req Request) *lua.LTable {
	lrequest := L.NewTable()
	lrequest.RawSetString("id", lua.LString(req.ID))
	lrequest.RawSetString("method", lua.LString(req.Method))
	lrequest.RawSetString("url", lua.LString(req.URL.String()))
	lrequest.RawSetString("path", lua.LString(req.URL.Path))
	lrequest.RawSetString("host", lua.LString(req.Host))
	lrequest.RawSetString("getPathParam", L.NewFunction(req.getPathParam))

	if req.Body == nil {
		lrequest.RawSetString("body", lua.LNil)
	} else {
		reqBody := L.NewTable()

		// Need to probably have users set Body Shape/struct/type via config or lua definition
		for k, v := range req.Body {
			reqBody.RawSetString(k, lua.LString(v.(string)))
		}
		lrequest.RawSetString("body", reqBody)
	}

	// Query Params
	lqueryParams := L.NewTable()
	for k, v := range req.URL.Query() {
		lqueryParams.RawSetString(k, util.BuildStringArray(L, v))
	}

	lrequest.RawSetString("queryParams", lqueryParams)

	// Headers
	lheaders := L.NewTable()
	for k, v := range req.Header {
		lheaders.RawSetString(k, util.BuildStringArray(L, v))
	}

	lrequest.RawSetString("headers", lheaders)

	return lrequest
}

// MakeLResponse binds the native GoLang sabaresu request object into the Lua state
func MakeLResponse(L *lua.LState, resp Response) *lua.LTable {
	lresponse := L.NewTable()
	lresponse.RawSetString("statusCode", lua.LNumber(200))
	lresponse.RawSetString("headers", L.NewTable())
	lresponse.RawSetString("payload", L.NewTable())
	lresponse.RawSetString("body", L.NewTable())

	return lresponse
}
