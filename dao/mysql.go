package dao

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:SHANG6688@tcp(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) // DB已经声明过了
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
