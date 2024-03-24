package lua

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Phamiliarize/sabaresu/pkg/gateway"
	"github.com/Phamiliarize/sabaresu/pkg/runtime/lua/util"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

type luaRuntime struct{}

func NewLuaRuntime() *luaRuntime {
	return &luaRuntime{}
}

func (l *luaRuntime) RuntimeHandler(funcDir string, functions []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		L := lua.NewState()
		defer L.Close()

		reqBody, err := util.Body(r)
		if err != nil {
			panic(err)
		}

		lreq := MakeLRequest(L, Request{
			ID:         r.Context().Value(gateway.RequestIDContextKey).(string),
			Method:     r.Method,
			URL:        r.URL,
			Header:     r.Header,
			Host:       r.Host,
			gPathParam: r.PathValue,
			Body:       reqBody,
		})

		lresp := MakeLResponse(L, Response{})

		// Chain through the Lua functions
		for i, function := range functions {
			// Set previous functions response if present
			if i > 0 {
				lresp = L.Get(-1).(*lua.LTable)
				L.Pop(1)
			}

			filePath := fmt.Sprintf("%s/%s", funcDir, function)

			if err := L.DoFile(filePath); err != nil {
				panic(err)
			}

			if err := L.CallByParam(lua.P{
				Fn:      L.GetGlobal("main"),
				NRet:    1,
				Protect: true,
			}, lreq, lresp); err != nil {
				panic(err)
			}

			// Set function response if final
			if i+1 == len(functions) {
				lresp = L.Get(-1).(*lua.LTable)
				L.Pop(1)
			}
		}

		var resp Response
		if err := gluamapper.Map(lresp, &resp); err != nil {
			panic(err)
		}

		res, err := json.Marshal(resp.Body)
		if err != nil {
			// TODO: need to handle errors not panic on everything
			panic(err)
		}

		w.WriteHeader(resp.StatusCode)
		// Only JSON is supported right now
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
