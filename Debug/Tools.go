package Debug

import (
	"runtime"
	"fmt"
	"strings"
)

/**
	返回调用位置信息
 */
func ErrorMsg(entity error) {
	if pc, file, line, ok := runtime.Caller(1); ok != false {
		f := runtime.FuncForPC(pc)
		fmt.Printf(" [调用方法] %s \n [调用位置] %s %v \n [ErrorMsg] %s\n", f.Name(), file, line, entity)
	}
	return
}

func RenderLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func RenderServer(format string, args ...interface{}){
	RenderLog("Server", 0, format, args...)
}

func RenderClient(format string, args ...interface{}){
	RenderLog("Client", 0, format, args...)
}
