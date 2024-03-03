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

#### The Function Chain

You can chain functions in `sabaresu`. This is done by simply writing the functions in an array, in the order you want them to fire.

Each function passes the sabaresu context object (which contains `request` and `response`) to the next function.

This allows a very expressive an easy way to think about the "flow" of a given request and also isolate parts to make testing easier.
