package cli

import (
	"fmt"
	"net/http"

	gw "github.com/Phamiliarize/sabaresu/pkg/gateway"
	"github.com/Phamiliarize/sabaresu/pkg/runtime/lua"
)

func Run(cfgPath *string, schemaPath *string, port *int) {
	gateway, err := gw.NewGatewayFromCfg(cfgPath, schemaPath)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	middleware := []gw.Middleware{gw.PanicRecovery, gw.RequestLogging, gw.RequestValidator, gw.ResponseSweeper}

	luaRuntime := lua.NewLuaRuntime()

	for _, r := range gateway.Routes {
		mux.HandleFunc(
			fmt.Sprintf("%s %s", r.Method, r.Path),
			gw.RegisterRuntimeMiddleware(middleware, luaRuntime.RuntimeHandler(r.FuncDir, r.Functions)),
		)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
}
