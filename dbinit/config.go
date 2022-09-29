package dbinit

import (
	"log"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

var (
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
)

func init() {
	workdir, err := os.Getwd()
	wrpath := "/dbinit/"
	var str = []string{workdir, wrpath, "config.ini"}
	path := strings.Join(str, "")

	config, err := ini.Load(path)
	checkErr(err)
	LoadMysqlData(config)
	dbConfig := []string{DbUser, ":", DbPassword, "@tcp(", DbHost, ")/", DbName, "?charset=utf8&parseTime=true"}
	pathDB := strings.Join(dbConfig, "")
	initDB(pathDB)

}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
