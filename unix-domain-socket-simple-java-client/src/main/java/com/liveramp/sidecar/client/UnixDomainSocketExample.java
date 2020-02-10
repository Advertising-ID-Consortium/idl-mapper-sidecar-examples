package com.liveramp.sidecar.client;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import org.newsclub.net.unix.AFUNIXSocket;
import org.newsclub.net.unix.AFUNIXSocketAddress;

public class UnixDomainSocketExample {

  private static final String LIVERAMP_IDL_MAPPER_SOCKET = "/tmp/liveramp/idl-mapper-socket";
  private static final String LIVERAMP_IDL_MAPPER_MESSAGE = "GET /map?env=AnZGl_biYwUNkglgfKZiXMLv1YEYkzTFzkgTIQ17F_i3l4ruMEzWNJDDm1Ofv7c0ynB51uFQO9ytmMGuDX8N8I89s2s HTTP/1.1\n"
      + "Host: localhost\n\n";

  public static void main(String[] args) throws IOException {

    File socketFile = new File(LIVERAMP_IDL_MAPPER_SOCKET);

    try (AFUNIXSocket sock = AFUNIXSocket.newInstance()) {
      sock.connect(new AFUNIXSocketAddress(socketFile));

      try (InputStream is = sock.getInputStream(); //
          OutputStream os = sock.getOutputStream();) {

        os.write(LIVERAMP_IDL_MAPPER_MESSAGE.getBytes(StandardCharsets.UTF_8));
        os.flush();
        
        int bytes;
        byte[] readBuf = new byte[sock.getReceiveBufferSize()];

        while ((bytes = is.read(readBuf)) != -1) {
          System.out.write(readBuf, 0, bytes);
          System.out.flush();
        }
      }
    }
  }
  
}
