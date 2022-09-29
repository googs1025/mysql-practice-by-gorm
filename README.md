# mysql-practice-by-gorm
## 利用gorm对mysql调用，的练习。
## 并使用pipeline管道模式，进行并发写入操作提升性能。

###目录结构
```bigquery
├── README.md   
├── csv # 存放csv结果的文件
│   └── 1.csv
├── dbinit  # db初始化与配置文件
│   ├── config.go
│   ├── config.ini
│   └── mysqlinit.go
├── getData # 从mysql中取数据，主逻辑
│   ├── getdata.go  # 一般模式
│   ├── getdata_by_pipeline_mode.go # 管道模式
│   └── readme.md
├── go.mod
├── go.sum
├── main.go
└── pipeline    # 抽象出来的pipeline框架
    ├── data.go
    ├── pipeline.go
    └── test.go



```
## 一般模式
![项目架构](https://github.com/googs1025/mysql-practice-by-gorm/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE%20(1).jpg?raw=true)
## pipeline管道模式

![项目架构](https://github.com/googs1025/mysql-practice-by-gorm/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE.jpg?raw=true)



