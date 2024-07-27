package base

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// contextValue 每执行一次withValue都会创建一个新的context, 所以contextValue每次都会返回一个新的context,
// 新的context 在形式上会覆盖旧的key,但是实际上，旧的key没有覆盖，只不过查找的时候会优先从最先创建的context中查找。
func contextValue() context.Context {
	begin := context.Background()
	ctx := context.WithValue(begin, "key1", "value1")
	// 此时key1的值被覆盖了（形式上）
	ctx = context.WithValue(ctx, "key1", "value2")
	// 可以创建多个key value映射，实际上创建了多个context
	ctx = context.WithValue(ctx, "key2", "value3")
	return ctx
}

func TestContextValue(t *testing.T) {
	ctx := contextValue()
	assert.Equal(t, "value2", ctx.Value("key1"))
	assert.Equal(t, "value3", ctx.Value("key2"))
}

func TestContextCancel(t *testing.T) {
	begin := context.Background()
	ctx, cancel := context.WithCancel(begin)

	go func() {
		select {
		case <-ctx.Done():
			t.Log("done")
			return
		}
	}()
	time.Sleep(time.Second * 10)
	cancel()
	// 确保cancel 之后等协程接收信号之后再退出主程
	time.Sleep(time.Second * 2)
}

func TestContextTimeout(t *testing.T) {
	// 2秒之后 <-ctx.Done() 有数据，也可以主动调用cancel取消
	begin := context.Background()
	ctx, cancel := context.WithTimeout(begin, time.Second*2)
	select {
	case <-ctx.Done():
		t.Log("done")
	}
	defer cancel()
}

func TestContextDeadline(t *testing.T) {
	// 2秒之后 <-ctx.Done() 有数据，也可以主动调用cancel取消
	begin := context.Background()
	ctx, cancel := context.WithDeadline(begin, time.Now().Add(time.Second*2))
	select {
	case <-ctx.Done():
		t.Log("done")
	}
	defer cancel()
}

// TestContextDeadlineCause 创建一个带有超时和取消原因的context
func TestContextDeadlineCause(t *testing.T) {
	// 2秒之后 <-ctx.Done() 有数据，也可以主动调用cancel取消
	begin := context.Background()
	ctx, cancel := context.WithDeadlineCause(begin, time.Now().Add(time.Second*2), errors.New("2秒超时"))
	select {
	case <-ctx.Done():
		t.Log(context.Cause(ctx))
		t.Log("done")
	}
	defer cancel()
}

func TestContextCancelAfterFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	// 建立 AfterFunc 和 ctx 的关联
	stop := context.AfterFunc(ctx, func() {
		t.Log("AfterFunc running")
	})
	// 停止 AfterFunc 和 ctx 的关联
	if stop() {
		t.Log("取消 AfterFunc 和parent 的关联成功")
	} else {
		t.Log("取消 AfterFunc 和parent 的关联失败")
	}
	time.Sleep(time.Second * 4)
}

/*
*

		=== RUN   TestContextCancelCause
	    context_test.go:107: 主动取消
		--- PASS: TestContextCancelCause (0.00s)
		PASS
*/
func TestContextCancelCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	ctx, _ = context.WithTimeoutCause(ctx, time.Second*2, errors.New("2秒超时"))
	cancel(errors.New("主动取消"))

	// 获取真正的原因
	t.Log(context.Cause(ctx))
}
