# IDL Mapper Sidecar GO Lang TCP Client Example


## About

This example shows how to make http requests to Liveramp's IDL Mapper Sidecar's `/health` and `/map` endpoints using Go lang. Example is done using standard "net/http" Go library.

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
> GET /health HTTP/1.1
> Host: http://localhost:3000
>
< HTTP/1.1 200 OK
< Content-Length: 0

> GET /map?env=AjfowUURXDJnQmc_HNeuswelMv4ZHZQJFM8TpiUnYEyA81Vdgg HTTP/1.1
> Host: http://localhost:3000
>
< HTTP/1.1 200 OK
< Content-Type: application/json
<
{"SEAT0040":"Xm0040hok0YnyQ2rXg1O_-7O6yG9SYW6-qYM4C_VBNzKK4ovc",
"SEAT0041":"Xm004162s7uX8XEBe72TSKI1zx1Ldys9HjcRDLy20W88_6fJY",
"SEAT0043":"Xm0043GG4d-qZjutb_XKwx-TexpjvSEY29hNqwRpfvt0r8Bo0",
.
.
.
"SEAT0009":"Xm0009muLqkYEpfnIJFbk22-aNRQ3THFxsvK7x-vW61CRn5hE",
"SEAT0017":"Xm0017Ho8ChiYxgA3tOWCvaAXHs7XSoC3usJvprCrJ4f9FnXE",
"SEAT0002":"Xm0002tusbcMNX8E5LuzcYrfIsEAEHB2jLVrhgVRezowJdmQs"}
```