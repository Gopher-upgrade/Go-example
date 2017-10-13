package main

import (
	"net"
	"time"
	"Go-example/Debug"
	"bytes"
	"bufio"
	"os"
	"strings"
	"fmt"
)

const (
	RE = '\t'
)

func ClientInit(netWork, address string) {
	conn, err := net.DialTimeout(netWork, address, 2*time.Second)
	if err != nil {
		Debug.RenderClient("Dial Error: %s", err)
		return
	}
	defer conn.Close()
	Debug.RenderClient("Connected to server. (remote address: %s, local address: %s)",
		conn.RemoteAddr(), conn.LocalAddr())
	for {
		cmdReader := bufio.NewReader(os.Stdin)
		cmdStr, err := cmdReader.ReadString('\n')
		if err != nil{
			break
		}
		//这里把读取的数据后面的换行去掉，对于Mac是"\r"，Linux下面
		//是"\n"，Windows下面是"\r\n"，所以为了支持多平台，直接用
		//"\r\n"作为过滤字符
		cmdStr = strings.Trim(cmdStr, "\r\n")
		time.Sleep(200 * time.Millisecond)
		conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
		var buffer bytes.Buffer
		buffer.WriteString(cmdStr)
		var delimiter byte
		delimiter = '\t'
		buffer.WriteByte(delimiter)
		conn.Write(buffer.Bytes())
	}
}

func readClient(conn net.Conn, delimiter string) (string, error) {
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
	ClientInit("tcp", "127.0.0.1:9630")
}
