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
  private static final String LIVERAMP_IDL_MAPPER_HEALTH_MSG = "GET /health HTTP/1.1\nHost: localhost\n\n";
  private static final String LIVERAMP_IDL_MAPPER_MAP_MSG = "GET /map?env=AnZGl_biYwUNkglgfKZiXMLv1YEYkzTFzkgTIQ17F_i3l4ruMEzWNJDDm1Ofv7c0ynB51uFQO9ytmMGuDX8N8I89s2s HTTP/1.1\n"
      + "Host: localhost\n\n";
  

  public static void main(String[] args) throws IOException {

    File socketFile = new File(LIVERAMP_IDL_MAPPER_SOCKET);

    try (AFUNIXSocket sock = AFUNIXSocket.newInstance()) {
      sock.connect(new AFUNIXSocketAddress(socketFile));

      try (InputStream is = sock.getInputStream(); //
          OutputStream os = sock.getOutputStream();) {

        os.write(LIVERAMP_IDL_MAPPER_HEALTH_MSG.getBytes(StandardCharsets.UTF_8));
        os.flush();
        
        byte[] readBuf = new byte[sock.getReceiveBufferSize()];
        int bytes = is.read(readBuf);
        
        System.out.println("Response from Sidecar's '/health' endpoint:");
        System.out.println();
        System.out.println(new String(readBuf, 0, bytes, StandardCharsets.UTF_8));

        os.write(LIVERAMP_IDL_MAPPER_MAP_MSG.getBytes(StandardCharsets.UTF_8));
        os.flush();
        
        readBuf = new byte[sock.getReceiveBufferSize()];
        bytes = is.read(readBuf);
        
        System.out.println();
        System.out.println("Response from Sidecar's '/map' endpoint:");
        System.out.println();
        System.out.println(new String(readBuf, 0, bytes, StandardCharsets.UTF_8));
        
      }
    }
  }
  
}
