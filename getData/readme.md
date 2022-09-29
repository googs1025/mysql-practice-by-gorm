# 順序執行改为管道模式执行的总结

## 第一步：改为管道
入参：（管道连接起点）

```bigquery
// 输出的管道参数是自定义要传入的struct
type InChan chan *BookList
```

出参：（管道数据输出）

```bigquery
type OutChan chan *Result
```

结果集：需要接口管道output的结果

```
// 除了放error外，可以放入业务需要的字段参数
type Result struct{
	Page int	// 打印的csv以1000为一页
	Err error
}
```

## 第二步：定义管道命令类型
```bigquery
type DataCmd func() InChan
type DataPipeCmd  func(in InChan) OutChan
```

## 第三步：管道函数
```go
func Pipe(c1 DataCmd,cs ...DataPipeCmd) OutChan{
	in:=c1()
	out:=make(OutChan)
	wg:=sync.WaitGroup{}
	for _,c:=range cs{
		getChan:=c(in)
		wg.Add(1)
		go func(input OutChan) {
			defer wg.Done()
			for v:=range input{
				out<-v
			}
		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out

}
```
