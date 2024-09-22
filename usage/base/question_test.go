package base

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var b = func() int {
	fmt.Printf("b")
	return 0
}()

var count int

func increment() {
	count++
	fmt.Printf("count: %d", count)
}

func TestIncrement(t *testing.T) {
	increment()
}

// 声明一个常量 pi，并将其初始化为 3.14。
const pi = 3.14

// iota 是一个iota常量，iota 的值从 0 开始递增，每次递增 1。
const (
	Red = iota
	Blue
	Yellow
)

// 类型转换

func TestTypeConversion(t *testing.T) {
	x := 10.5
	y := int(x)
	fmt.Printf("x: %f, y: %d", x, y)
}

// 变量隐蔽问题
func TestVariableHiding(t *testing.T) {
	var a int
	var b int

	if b == 1 {
		a := 1
		assert.Equal(t, 1, a)
	} else {
		a := 2
		assert.Equal(t, 2, a)
	}

	// a 被隐蔽了，始终为 0
	assert.Equal(t, 0, a)
}
