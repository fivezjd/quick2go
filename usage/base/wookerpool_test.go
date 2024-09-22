package base

import (
	"errors"
	"fmt"
	e "github.com/pkg/errors"
	"testing"
	"time"
)

type worker struct {
	id  int
	err error
}

func (w *worker) run(ch chan *worker) {
	defer func() {
		if err := recover(); err != nil {
			w.err = fmt.Errorf("worker %d panicked: %v", w.id, err)
		} else {
			w.err = fmt.Errorf("worker %d panicked", w.id)
		}
		ch <- w
	}()

	fmt.Printf("%d is running", w.id)
	time.Sleep(2 * time.Second)
	panic(fmt.Errorf("worker %d panicked", w.id))
}

type WorkerPool struct {
	workerNum  int
	workerChan chan *worker
}

func NewWorkerPool(workerNum int) *WorkerPool {
	return &WorkerPool{
		workerNum:  workerNum,
		workerChan: make(chan *worker, workerNum),
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.workerNum; i++ {
		wk := &worker{id: i}
		go wk.run(wp.workerChan)
	}
}
func (wp *WorkerPool) Wait() {
	for wk := range wp.workerChan {
		fmt.Printf("%d is done, err: %v\n", wk.id, wk.err)
		wk.err = nil
		go wk.run(wp.workerChan)
	}
}

func TestWorkerPool(t *testing.T) {
	wp := NewWorkerPool(10)
	wp.Run()
	wp.Wait()
}

// 模拟函数，可能会返回错误
func doSomething() error {
	return errors.New("something went wrong")
}

// 处理错误并打印调用栈
func handleError(err error) {
	if err != nil {
		// 使用 WithStack 包装错误以附加调用栈
		errWithStack := e.WithStack(err)

		// 打印错误信息
		fmt.Printf("Error occurred: %v\n", errWithStack)

		// 尝试提取调用栈
		if stackErr, ok := errWithStack.(interface {
			StackTrace() e.StackTrace
		}); ok {
			stackTrace := stackErr.StackTrace()
			for _, frame := range stackTrace {
				// 打印调用栈的每一帧
				fmt.Println(frame)
			}
		} else {
			fmt.Println("No stack trace available")
		}
	}
}

func TestE(t *testing.T) {
	if err := doSomething(); err != nil {
		handleError(err)
	}
}
