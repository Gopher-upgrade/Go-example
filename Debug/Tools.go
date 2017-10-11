package Debug

import (
	"runtime"
	"fmt"
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
