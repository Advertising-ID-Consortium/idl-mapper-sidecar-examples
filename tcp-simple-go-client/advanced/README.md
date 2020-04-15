# IDL Mapper Sidecar GO Lang TCP Client Example


## About

This example shows how to make http requests to Liveramp's IDL Mapper Sidecar's `/health` and `/map` endpoints using Go lang. Example is done using `IDLMapperClient` that act as wrapper.

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
INFO[0000] IDL Mapper's '/health' endpoint responded with status code: 200 
INFO[0000] IDL Mapper's '/map' endpoint for envelope 'AjfowUURXDJnQmc_HNeuswelMv4ZHZQJFM8TpiUnYEyA81Vdgg' responded with mappings: map[SEAT0000:Xm00008qJ4ii7hBlZ8DX8WrQPKISWMT8AqwAWL3TMZkXcd080 SEAT0001:Xm0001mqTlQbEQHG42IMVVqxJt66OVxG6m3fobgmPWnzcHeyk . . .]
```