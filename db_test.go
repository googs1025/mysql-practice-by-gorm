package main

import (
	"fmt"
	"mysql-practice-by-gorm/dbinit"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type TestTable struct {
	Name string
	Age int

}

func TestMySqlDB(t *testing.T) {
	Convey("test db talbe", t, func() {
		// 创建表结构
		if !dbinit.DB.Migrator().HasTable(&TestTable{}) {
			_ = dbinit.DB.AutoMigrate(&TestTable{})
		}

		user := &TestTable{Name: "Jinzhu111", Age: 18111}

		// 插入数据。
		result := dbinit.DB.Create(user)
		So(result.Error, ShouldBeNil)
		fmt.Println(result.Error, result.RowsAffected)
		// 删除数据。
		result = dbinit.DB.Where("name=?", "Jinzhu111").Delete(user)
		fmt.Println(result.Error, result.RowsAffected)
		So(result.Error, ShouldBeNil)
		// 删除测试表。
		_ = dbinit.DB.Migrator().DropTable(user)
	})



}