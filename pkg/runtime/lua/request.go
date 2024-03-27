package lua

import (
	"net/http"
	"net/url"

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
