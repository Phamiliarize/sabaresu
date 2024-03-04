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
Functions are exposed over HTTP; you choose how and what to expose via `gateway.toml`, which follows a simple mental model:

```toml
[[routes]]
method = "GET" # HTTP Method
path = "/v1/test" # Path
functions = ["hello-world.lua"]

[[routes]]
method = "GET"
path = "/v1/test/{test}" # Path Parameter
functions = ["auth.lua", "hello-world.lua"] # Function Chain
```

### Functions

You might notice above that we listed functions in an array- this is becase you can chain funtions in `sabaresu`. The runtme will fire your functions in order- allowing for an easy method to isolate each step of your application flow.


#### I/O
Every time your Lua function fires it is given the original `request` table- this is immutable between functions.

Your function is expected to return a `response` table which is mutable between functions, and allows you to pass data forward and mutate the expected response. We can see how this works below:

```lua
function main(req, resp)
    print(req.requestId)
    print("Hello " .. req.getPathParam("name") .. "!")
    return resp
 end
```

After your final function runs the results of `response` will form your HTTP response.

## Runtime API

Every function is supplied a `request` and `response`. Some helper functions are also made available as part of our "runtime API".

#### `request`
| key | description |
| ---- | ---- |
| `requestId` | A unique `UUIDv4` per HTTP request |
| `method` | HTTP Method |
| `url` | The full URL string |
| `path` | The path |
| `host` | The host value |
| `getPathParam("name")` | Takes the name of an URL parameter; returns value or empty string |


#### `request`
| key | description | default |
| ---- | ---- | ---- |
| `statusCode` | A number representing the HTTP Status Code | 200 |
