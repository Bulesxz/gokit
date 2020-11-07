package gokit

import (
	"context"
	"fmt"
	"sync"
)

// Job 任务对象
type Job interface {
	Do()
}

// Worker 工作worker
type Worker struct {
	JobQueue chan Job
	ctx      context.Context
	wg       *sync.WaitGroup
}

// NewWorker new worker
func NewWorker(ctx context.Context, wg *sync.WaitGroup) Worker {
	return Worker{ctx: ctx, JobQueue: make(chan Job), wg: wg}
}

// Run worker working
func (w Worker) Run(wq chan chan Job) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			wq <- w.JobQueue // 报告worker可用
			select {
			case job := <-w.JobQueue:
				job.Do() //
			case <-w.ctx.Done():
				return
			}
		}
	}()
}

// WorkerPool worker pool
type WorkerPool struct {
	ctx         context.Context
	workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
	wg          *sync.WaitGroup
}

// NewWorkerPool new worker pool
func NewWorkerPool(ctx context.Context, workerlen int, wg *sync.WaitGroup) *WorkerPool {
	return &WorkerPool{
		ctx:         ctx,
		wg:          wg,
		workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

// Run worker pool working
func (wp *WorkerPool) Run() {
	fmt.Println("init worker")
	//初始化worker
	for i := 0; i < wp.workerlen; i++ {
		worker := NewWorker(wp.ctx, wp.wg)
		worker.Run(wp.WorkerQueue)
	}
	// 循环获取可用的worker,往worker中写job
	wp.wg.Add(1)
	go func() {
		defer wp.wg.Done()
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			case <-wp.ctx.Done():
				break
			}
		}
	}()
}

// AddJob add job to worker pool
func (wp *WorkerPool) AddJob(job Job) {
	wp.JobQueue <- job
}
