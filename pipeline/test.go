package pipeline


import (
	"fmt"
	"mysql-practice-by-gorm/dbinit"

	"time"
	"log"
)

const sql="select * from books order by book_id limit ? offset ?"

func GetPage(args ...interface{}) InChan  {
	in:=make(InChan)
	go func() {
		defer close(in)
		for i:=1;i<=80;i++{
			in<-i
		}
	}()
	return in


}
func GetData(in InChan) OutChan {
	out:=make(OutChan)
	go func() {
		defer close(out)
		for d:=range in {
			page:=d.(int)
			pagesize:=1000
			booklist := &BookList{make([]*Book, 0), page}
			db := dbinit.DB.Raw(sql, pagesize, (page-1)*pagesize).Find(&booklist.Data)
			if db.Error!=nil{
				log.Println(db.Error)
			}
			out<-booklist.Data
		}
	}()
	return out
}

//模拟处理数据
func DoData(in InChan) OutChan{
	out:=make(OutChan)
	go func() {
		defer close(out)
		for d:=range in {
			v:=d.([]*Book)
			time.Sleep(time.Second*1)
			out<-fmt.Sprintf("处理了%d条数据,%d\n",len(v),time.Now().Unix())
		}
	}()
	return out
}
func PipeTest()  {
	p1:=NewPipe()
	p1.SetCmd(GetPage)
	p1.SetPipeCmd(GetData,3)
	out:=p1.Exec()


	p2:=NewPipe()
	p2.SetCmd(func(args ...interface{}) InChan {
		return InChan(out)
	})
	p2.SetPipeCmd(DoData,2)
	out2:=p2.Exec()

	for item:=range out2{
		fmt.Println(item)
	}

}
