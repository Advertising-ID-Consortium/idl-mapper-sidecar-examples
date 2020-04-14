# IDL Mapper Sidecar GO Lang TCP Client Example


## About

This example shows how to make http requests to Liveramp's IDL Mapper Sidecar's `/health` and `/map` endpoints using Go lang. In example request to Sidecar are done using `IDLMapperClient` that act as wrapper and can be imported into your project.

## Run Liveramp's IDL Mapper Sidecar

The following example shows how to start the [Liveramp's IDL Mapper Sidecar](https://sidecar.readme.io) container ready for TCP communication.

```bash
docker run --cap-add IPC_LOCK -p 127.0.0.1:3000:3000 liveramp/idl-mapper:latest-dev -listen tcp://:3000 -certificate " " [other args]
```

If the previous command is executed successfully you should see log messages like these:

```bash
time="2020-04-13T19:12:03Z" level=info msg="Set up keepalive for TCP"
time="2020-04-13T19:12:03Z" level=info msg=serving address=":3000"
```

## Run Example

In current folder execute following:

```bash
go run .
```

Example app will make requests to `/health` and `/map` endpoint. Results will be displayed in following format:

```
INFO[0000] IDL Mapper's '/health' endpoint responded with status: 200 OK 
INFO[0000] IDL Mapper's '/map' endpoint for envelope 'AjfowUURXDJnQmc_HNeuswelMv4ZHZQJFM8TpiUnYEyA81Vdgg' responded with mappings: {"SEAT0023":"Xm0023J5H4LcNJYOdeECH7iFfl9Ib0WTUQfU,"SEAT0047":"Xm0047MkRGchQFgiqnzB96yjgdV0blW-VElytUqYVQGHpsX6c" . . .} 
```

## Import `IDLMapperClient` into your project


### Pre go.mod

Get package:
```bash
go get github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced
```

Import (as shown in `main.go`):
```
import "github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced/idlmapperclient"
```

Use client (as shown in `main.go`):
```
// get configured idl mapper client instance
idlMapperClient := idlmapperclient.GetIDLMapperClient("http://localhost", 3000)

// call health
resp, err := idlMapperClient.Health()
```


### Post go.mod

Add dependency to go.mod file and use replace directive: 
```
require github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced v0.0.0

replace github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced v0.0.0 => <path_to_downloaded_repo>/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced
```

Import (as shown in `main.go`):
```
import "github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced/idlmapperclient"
```

Use client (as shown in `main.go`):
```
// get configured idl mapper client instance
idlMapperClient := idlmapperclient.GetIDLMapperClient("http://localhost", 3000)

// call health
resp, err := idlMapperClient.Health()
```