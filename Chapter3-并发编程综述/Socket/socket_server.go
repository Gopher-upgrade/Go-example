package main

import (
	"net"
	"Go-example/Debug"
	"fmt"
	"bytes"
	"io"
	"strings"
)

func serverInit(netWork, address, delimiter string) {
	var listener net.Listener
	listener, err := net.Listen(netWork, address) // 1. 开始监听9630 端口
	if err != nil {
		Debug.ErrorMsg(err)
		return
	}
	defer listener.Close() // 2. 代码结束后释放监听资源
	fmt.Printf("建立连接成功 %s\n", listener.Addr())
	for {
		conn, err := listener.Accept() // 阻塞直至新连接到来。
		if err != nil {
			Debug.RenderServer("Accept Error: %s", err)
		}
		go func(conn net.Conn) { // 3. 连接建立成功
			defer func() { // 4. 连接结束释放资源
				conn.Close()
				//wg.Done()
			}()
			client := strings.Replace(conn.RemoteAddr().String(), "127.0.0.1:", "", 3)
			Debug.RenderServer("%s 加入进来了", client)
			for {
				//conn.SetReadDeadline(time.Now().Add(30 * time.Second)) // 设置读取长
				strReq, err := readServer(conn, delimiter)
				if err != nil {
					if err == io.EOF {
						Debug.RenderServer(" %s 断开了连接", client)
					} else {
						Debug.RenderServer("Read Error: %s", err)
					}
					break
				}
				Debug.RenderServer("%s 说：%s", client, strReq)
			}
		}(conn)
	}
}

func readServer(conn net.Conn, delimiter string) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if string(readByte) == delimiter {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func main() {
	serverInit("tcp", "127.0.0.1:9630", "\t")
}
