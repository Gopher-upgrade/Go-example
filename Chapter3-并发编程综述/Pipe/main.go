/**
 * Pipe 管道学习
 *
 * @author   ShaoWei Pu  <marco0727@gamil.com>
 * @date     2017/10/9
 * -------------------------------------------------------------
 * 0x01. exec包执行外部命令，它将os.StartProcess进行包装使得它更容易映射到stdin和stdout，并且利用pipe连接i/o．
 * 0x02. 命名管道可以被多路复用
 */

package main

import (
	"os/exec"
	"fmt"
	"Go-example/Debug"
	"bytes"
	"time"
	"io"
	"os"
)

// output echo
func SimpleEcho() {
	cmd := exec.Command("echo", "-n", "I'm Echo")
	if err := cmd.Start(); err != nil {
		Debug.ErrorMsg(err)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Debug.ErrorMsg(err)
		return
	}
	// 保证关闭输出流
	defer stdout.Close()
	output := make([]byte, 30)
	n, err := stdout.Read(output)
	if err != nil {
		Debug.ErrorMsg(err)
		return
	}
	fmt.Printf("%s\n", output[:n])
}

// output pipe
func anonymousPipe() {
	cmdLeft := exec.Command("ps", "aux")
	cmdRight := exec.Command("grep", "php")
	var outputBuf bytes.Buffer
	cmdLeft.Stdout = &outputBuf // stdout
	if err := cmdLeft.Start(); err != nil {
		Debug.ErrorMsg(err)
		return
	}
	if err := cmdLeft.Wait(); err != nil {
		Debug.ErrorMsg(err)
		return
	}
	cmdRight.Stdin = &outputBuf // stdin
	var outputBufRight bytes.Buffer
	cmdRight.Stdout = &outputBufRight
	if err := cmdRight.Start(); err != nil {
		Debug.ErrorMsg(err)
		return
	}
	if err := cmdRight.Wait(); err != nil {
		Debug.ErrorMsg(err)
		return
	}
	fmt.Printf("%s \n", outputBufRight.Bytes())
}

// 原子操作
func fileBasedPipe() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("Error: Couldn't create the named pipe: %s\n", err)
	}
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 0; i < 26; i++ {
		input[i] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

func main() {
	// SimpleEcho() // 通过调用Linux命令输出Echo // Ps： 妈的 第一个Demo 就不是想象中的样子   exec: StdoutPipe after process started
	// anonymousPipe() // ps aux | grep php << 这个好使
	fileBasedPipe()
	inMemorySyncPipe()
}
