package syntaxBase

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

type poolTask interface {
	Run()
	Stop()
	IsRunning() bool
	//IsStopped() bool
	//GetResult() interface{}
}

type Pool struct {
	tasks   chan poolTask
	itemNum int
}

func (p *Pool) AddTask(task poolTask) {
	p.tasks <- task
}

func (p *Pool) Start() {
	for i := 0; i < p.itemNum; i++ {
		go func() {
			for {
				select {
				case task := <-p.tasks:
					task.Run()
				}
			}
		}()
	}
}

type option func(*Pool)

func NewPool(options ...option) *Pool {
	itemNum := runtime.NumCPU()
	pool := &Pool{
		tasks:   make(chan poolTask, itemNum),
		itemNum: itemNum,
	}
	for _, opt := range options {
		opt(pool)
	}
	return pool
}

type TaskOne struct {
	isRunning bool
	//isStopped bool
}

func (t *TaskOne) Run() {
	t.isRunning = true
}

func (t *TaskOne) IsRunning() bool {
	return t.isRunning
}

func (t *TaskOne) Stop() {
	t.isRunning = false
}

type TaskTwo struct {
	isRunning bool
}

func (t *TaskTwo) Run() {
	t.isRunning = true
}

func (t *TaskTwo) IsRunning() bool {
	return t.isRunning
}

func (t *TaskTwo) Stop() {
	t.isRunning = false
}

func TestGoroutine(t *testing.T) {
	pool := NewPool(
		func(p *Pool) {
			p.itemNum = 10
		},
	)
	assert.Equal(t, 10, pool.itemNum)
	pool.Start()

	taskOne := &TaskOne{}
	pool.AddTask(taskOne)

	taskTwo := &TaskTwo{}
	pool.AddTask(taskTwo)

	time.Sleep(time.Second * 2)

	// todo

	assert.Equal(t, true, taskOne.IsRunning())

	assert.Equal(t, true, taskTwo.IsRunning())

	taskOne.Stop()
	assert.Equal(t, false, taskOne.IsRunning())

	taskTwo.Stop()
	assert.Equal(t, false, taskTwo.IsRunning())

}
