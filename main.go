package main

import (
	pipeline "mysql-practice-by-gorm/pipeline"
)

func main()  {
	//getData.ReadData()
	//getData.Test()
	pipeline.PipeTest()
}

// TODO: 实现基于docker的mysql主从读写分离 (https://www.jianshu.com/p/ab20e835a73f)，dao层需要修改 读操作执行一个数据库 写操作执行一个数据库
