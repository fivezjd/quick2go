package syntaxBase

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	runtime.GOMAXPROCS(1)

	go func() {
		fmt.Println("子协程 运行中")
	}()
	// 让出CPU时间片
	runtime.Gosched()
}
