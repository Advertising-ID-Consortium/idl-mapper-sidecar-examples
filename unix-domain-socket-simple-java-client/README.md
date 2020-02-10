DL Mapper Sidecar JAVA Unix Domain Socket Client Example

[Liveramp's IDL Mapper Sidecar](https://sidecar.readme.io) container can communicate with your application via a Unix socket. The client will exchange HTTP messages over Unix Domain Socket. 
In order to do that, you will need to configure Sidecar to listen on a socket mounted inside Sidecar container. 
The following example shows how to start the Sidecar container with a Unix socket using a `/tmp/sidecar` folder mounted for Unix socket communication.

```bash
docker run --cap-add IPC_LOCK -v /tmp/liveramp:/mnt/sock liveramp/idl-mapper:latest-dev -listen unix:///mnt/sock/idl-mapper-socket -certificate " " [other args]
```
> **Note:**  This example needs to be run either on a Linux VM or in a native Linux environment.

If the previous command is executed successfully you should see a message like:
```bash
time="2020-02-05T12:03:39Z" level=info msg=serving address=/mnt/sock/idl-mapper-socket
```

This is a simple JAVA client to demonstrate communication with the Liveramp's IDL Mapper Sidecar via Unix Domain Socket.


>**Note:** This example is using [ junixsocket](https://github.com/kohlschutter/junixsocket) library.

