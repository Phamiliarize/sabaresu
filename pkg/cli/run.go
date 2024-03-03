package cli

import (
	"fmt"
	"net/http"

	"github.com/Phamiliarize/sabaresu/pkg/configuration"
	"github.com/Phamiliarize/sabaresu/pkg/lruntime"
)

func Run() {
	// Load Configuration
	routes, err := configuration.LoadGatewayConfiguration()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	for _, r := range *routes {
		mux.HandleFunc(fmt.Sprintf("%s %s", r.Method, r.Path), lruntime.RuntimeHandler(r.Functions))
	}

	http.ListenAndServe(":3000", mux)
}
