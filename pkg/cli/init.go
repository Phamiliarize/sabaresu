package cli

import (
	"log"
	"os"
)

func Init() {
	d1 := []byte(`[[routes]]
method = "GET"
path = "/v1/test"
functions = ["hello-world.lua"]`)

	d2 := []byte(`print("hello world")`)

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
