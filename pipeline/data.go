package pipeline

import "fmt"

// 数据库表对象
type Book struct {
	BookId int `gorm:"column:book_id"`
	BookName string `gorm:"column:book_name"`
}

// 需要的方法
func(this *Book) String() string{
	return fmt.Sprintf("bookid:%d,book_name:%s\n",this.BookId,this.BookName)
}

// list对象 + 分页
type BookList struct {
	Data []*Book
	Page int
}

// 结果
type Result struct{
	Page int
	Err error
}
