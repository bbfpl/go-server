package library

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	config := NewConfig("")
	dbType := config.Get("db.type")
	//dbHostname := config.Get("db.hostname")
	//dbHostport := config.Get("db.hostport")
	dbDatabase := config.Get("db.database")
	dbUsername := config.Get("db.username")
	dbPassword := config.Get("db.password")

	link := dbUsername.(string) + ":" + dbPassword.(string) + "@/" + dbDatabase.(string) + "?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(dbType.(string), link)

	if err != nil {
		fmt.Printf("mysql 连接错误 %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("数据库错误 %v", DB.Error)
	}
}
