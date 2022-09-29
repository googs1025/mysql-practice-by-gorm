package getData

import (
	"encoding/csv"
	"fmt"
	"mysql-practice-by-gorm/dbinit"
	"os"
	"strconv"
	"sync"
	"time"
)

type Book struct {
	BookId int `gorm:"column:book_id"`
	BookName string `gorm:"column:book_name"`
}
type BookList struct {
	Data []*Book
	Page int
}

// 存放结果struct
type Result struct{
	Page int	// 打印的csv以1000为一页
	Err error
}

// 入参(管道连接点)
type InChan chan *BookList
// 出参(管道最后输出)
type OutChan chan *Result


type DataCmd func() InChan
type DataPipeCmd  func(in InChan) OutChan

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

const sql="select * from books order by book_id limit ? offset ? "
func ReadData() InChan  {
	page:=1
	pagesize:=1000
	in:=make(InChan)
	go func() {
		defer close(in)
		for {
			booklist := &BookList{make([]*Book, 0), page}
			db := dbinit.DB.Raw(sql, pagesize, (page-1)*pagesize).Find(&booklist.Data)
			if db.Error != nil || db.RowsAffected == 0 {
				break
			}
			in<-booklist
			page++

		}
	}()
	return in


}

//管道函数
func WriteData(in InChan) OutChan  {
	out:=make(OutChan)
	go func() {
		defer close(out)
		for d:=range in {
			out<-&Result{Page:d.Page,Err:SaveData(d)}
		}
	}()
	return out
}


//写入到csv文件
func SaveData(data *BookList) error   {
	time.Sleep(time.Millisecond*500)
	file:=fmt.Sprintf("./csv/%d.csv",data.Page)
	csvFile,err:= os.OpenFile(file,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	if err!=nil{
		return err
	}
	defer csvFile.Close()
	w := csv.NewWriter(csvFile)//创建一个新的写入文件流
	header := []string{"book_id", "book_name"}
	export := [][]string{
		header,
	}
	for _,d:=range data.Data{
		cnt:=[]string{
			strconv.Itoa(d.BookId),
			d.BookName,
		}
		export=append(export,cnt)
	}
	err=w.WriteAll(export)
	if err!=nil{
		return err
	}
	w.Flush()
	return nil
}

func Test()  {
	out:=Pipe(ReadData,WriteData,WriteData,WriteData,WriteData,WriteData)
	for o:=range out{
		fmt.Printf("%d.csv文件执行完成,结果:%v\n",o.Page,o.Err)
	}
}

