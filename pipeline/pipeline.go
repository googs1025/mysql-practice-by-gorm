package pipeline

import "sync"

// pipeline框架

//
type InChan chan interface{}
type OutChan chan interface{}
type CmdFunc func(args ...interface{}) InChan
type PipeCmdFunc func(in InChan) OutChan

// 管道模式对象
type Pipe struct{
	Cmd CmdFunc	// 某个方法操作，方法output是管道，用于输入给下游多个管道用
	PipeCmd PipeCmdFunc	// 下游方法操作，多是耗时操作，需要用多个goroutine进行。
	Count int	// 多路复用的路有多少个
}

func NewPipe()  *Pipe {
	return &Pipe{Count:1}
}

func(p *Pipe) SetCmd(c CmdFunc)  {
	p.Cmd=c
}
func(p *Pipe) SetPipeCmd(c PipeCmdFunc,count int )  {
	p.PipeCmd=c
	p.Count=count
}

// Exec 执行函数
func(p *Pipe) Exec(args ...interface{}) OutChan  {
	in := p.Cmd(args)
	out := make(OutChan)
	wg := sync.WaitGroup{}
	for i:=0 ; i< p.Count; i++ {
		getChan := p.PipeCmd(in)
		wg.Add(1)
		go func(input OutChan) {
			defer wg.Done()
			for v := range input{
				out <-v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}
