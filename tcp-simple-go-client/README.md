# IDL Mapper Sidecar GO Lang TCP Client Example


## About

These examples shows how to call Liveramp's IDL Mapper Sidecar's `/health` and `/map` endpoints using Go lang. In [simple](https://github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tree/master/tcp-simple-go-client/simple) example, request are made manually using standard 'net/http' go library. In [advanced](https://github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tree/master/tcp-simple-go-client/advanced) example, request are made using `IDLMapperClient` that act as wrapper and can be imported and used in your project.