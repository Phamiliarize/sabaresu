package cli

import (
	"log"
	"os"
)

func Init() {
	d1 := []byte(`[[routes]]
method = "POST"
path = "/v1/enterprise/{enterprise_id}"
schema = "newEmployee"
funcDir = "./functions"
functions = ["example.lua"]`)

	d2 := []byte(`function main(request, response)
    response.body["name"] = request.getPathParam("name")
    return response
end`)

	d3 := []byte(`[path] # Validation of path params
enterprise_id = "string,required,uuid"

[body] # Validation of request body
name = "string,required,min=1,max=128"

[body.org_info]
id = "string,required,uuid"
name = "string,required,min=1,max=128"
salary = "number,required"

[response] # Definition of the response
id = "string,required,uuid"
enterprise_id = "string,required,uuid"
name = "string,required,min=1,max=128"

[response.org_info]
id = "string,required,uuid"
name = "string,required,min=1,max=128"`)

	// Write the gateway file
	if err := os.WriteFile("./gateway.toml", d1, 0644); err != nil {
		log.Fatal(err)
	}

	// Setup a functions folder
	if err := os.Mkdir("./functions", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Setup a schema folder
	if err := os.Mkdir("./schema", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./schema/newEmployee.toml", d3, 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./functions/example.lua", d2, 0644); err != nil {
		log.Fatal(err)
	}
}
