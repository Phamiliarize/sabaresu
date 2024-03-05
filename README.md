# 🐟 sabaresu (WIP)
No mackerels needed. Or was it servers!?

### Okay what is it really? Why?
sabaresu is a "toy" serverless framework built from my love for the DX of serverles but my desire to not be stuck to some FaaS platform.

Short-Term Goals:
- Scale easy: deploys like any other containerized-GoLang Web App
- Wizard DX:
    - Up and running in < 1 Minute
    - Easy to test functions

Moonshot Goals:
- Complex Function Deployments
    - Blue-Green
    - Canary
    - Hot-Changes
- No Fail Mode: auto fail deployment attempts if a function fails a test


## Getting Started
Build `sabaresu`:
```shell
go build -o tmp/sabaresu cmd/sabaresu/main.go
```

Initialize a new project:
```shell
cd tmp && ./sabaresu init
```

Run your server:
```shell
./sabaresu run
```



### `gateway.toml`
> Currently **only JSON** apis are is supported.

Functions are exposed over HTTP; you choose how and what to expose via `gateway.toml`, which follows a simple mental model:

```toml
[[routes]]
method = "GET" # HTTP Method
path = "/v1/test" # Path
functions = ["hello-world.lua"]

[[routes]]
method = "GET"
path = "/v1/user/{test}" # Path Parameter
functions = ["auth.lua", "hello-world.lua"] # Chaining functions auth -> hello-world
```

Chaining can be a powerful ally to isolating your application logic. You can treat them like befor-hooks or after-hooks. Middlewares. Whatever it needs to be.

### Functions

A sabaresu function must have define a `main` function that recieves two parameters, `request` & `response`. For example:

```lua
function main(request, response)
    print(request.id)
    print("Hello " .. request.getPathParam("name") .. "!")
    return resp
 end
```


#### request
`request` is a table that describes the original HTTP request recieved by the gateway.  Immutable between requests.

| key | type | description |
| ---- | ---- | ---- |
| `id` | string | A unique `UUIDv4` for the the request |
| `method` | string | The HTTP Method of the request |
| `headers` | table[string][]string | A table of headers on the request |
| `url` | string | The full URL|
| `host` | string | The host value |
| `path` | string | The request path |
| `getPathParam ` | function(name string) string or nil | Retrieves a path parameter by name |
| `queryParams` | table[string][]string | A table of query params on the request |


#### response
`response` is a table that defines how the final response will be sent after the last function is run. It is mutable between functions.

| key | type | description | default |
| ---- | ---- | ---- |  ---- |
| `statusCode` | number | HTTP Status Code | 200 |
| `headers` | table[string][]string | Headers to set on response | {} |
| `payload` | table[string]any | A private payload to pass info between functions| {}
| `body` | table[string]any | The response body that will be serialized | {} |


## Runtime API

In addition to request handling basics, some helpers and go bindings are also globaly available for usage.

