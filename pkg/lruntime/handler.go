package lruntime

import (
	"fmt"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

// RuntimeHandler generates a HandlerFunc from provided config
func RuntimeHandler(functions []string) http.HandlerFunc {
	middleware := []Middleware{PanicRecovery}

	return RegisterRuntimeMiddleware(middleware, func(w http.ResponseWriter, r *http.Request) {
		// SetSabaresu(r, "hi")

		// fmt.Println(r.PathValue("name"))

		// Chain through the Lua functions
		for _, function := range functions {

			filePath := fmt.Sprintf("./functions/%s", function)
			// SetSabaresu(r, function)

			L := lua.NewState()
			defer L.Close()
			if err := L.DoFile(filePath); err != nil {
				panic(err)
			}
		}

		fmt.Fprintf(w, "%s\n", GetSabaresu(r))
	})
}
