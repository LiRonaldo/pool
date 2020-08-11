package main

import (
	"fmt"
	"time"
)

//定义任务
type task struct {
	//任务要执行的方法。
	f func() error
}

//创建一个task
func NewTask(arg_f func() error) *task {
	t := task{f: arg_f}
	return &t
}

//执行任务里的方法
func (t task) Excuted() {
	t.f()
}

//定义一个线程池
type pool struct {
	EntityChannal chan *task
	JobsChannal   chan *task
	Worker_num    int
}

func NewPool(num int) *pool {
	p := pool{EntityChannal: make(chan *task),
		JobsChannal: make(chan *task),
		Worker_num:  num}
	return &p
}

//协成池创建worker
func (p *pool) worker(workId int) {
	//每次都从jobchannal中拿task，执行
	for task := range p.JobsChannal {
		task.Excuted()
		fmt.Println("Worker ID", workId, "执行代码")
	}
}
func (p *pool) run() {
	//创建worker
	for i := 0; i < p.Worker_num; i++ {
		go p.worker(i)
	}
	//从入口job中拿task放到jobchannal中
	for task := range p.EntityChannal {
		p.JobsChannal <- task
	}
}
func main() {
	//创建task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	p := NewPool(3)
	//开启一个协成，每次都放task入口jobchannal
	go func() {
		for {
			p.EntityChannal <- t
		}
	}()
	p.run()
}
