package lua

import (
	"github.com/Phamiliarize/sabaresu/pkg/runtime/lua/util"
	lua "github.com/yuin/gopher-lua"
)

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
