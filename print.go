package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func PrintMsg(msg string, strs ...string) {
	if len(strs) == 0 {
		fmt.Printf("[%s][info][%s][%s]\n", time.Now().Format(FormatTime), FuncCaller(), Green(msg))
		return
	}
	fmt.Printf("[%s][%s][%s][%s]\n", time.Now().Format(FormatTime), strs[0], FuncCaller(), Green(msg))
}

func PrintFatalError(err error) {
	fmt.Printf("[%s][error][%s][%s]\n", time.Now().Format(FormatTime), FuncCaller(), Red(err.Error()))
	os.Exit(1)
}

func FuncCaller() string {
	funcName, file, line, ok := runtime.Caller(2)
	filePaths := strings.Split(file, "/")
	if len(filePaths) > 2 {
		file = strings.Join(filePaths[len(filePaths)-2:], "/")
	}
	funcPaths := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	if ok {
		return fmt.Sprintf("caller:%s:%d func:%s", file, line, funcPaths[len(funcPaths)-1])
	}
	return ""
}
