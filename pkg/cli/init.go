package cli

import (
	"log"
	"os"
)

func Init() {
	d1 := []byte(`[[routes]]
method = "GET"
path = "/v1/test/{name}"
functions = ["hello-world.lua"]`)

	d2 := []byte(`function main(req, resp)
    print(req.requestId)
    print("Hello " .. req.getPathParam("name") .. "!")
    return resp
 end`)

	if err := os.WriteFile("./gateway.toml", d1, 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir("./functions", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./functions/hello-world.lua", d2, 0644); err != nil {
		log.Fatal(err)
	}
}
