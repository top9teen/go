package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB
var err error

func Init() {

	db, err = gorm.Open("mssql", "sqlserver://"+USER+":"+PASSWORD+"@"+HOST+":"+PORT+"?database="+DB+"")

	db.LogMode(true)

	if err != nil {
		panic("connectFail")
	}

}

func DBsql() *gorm.DB {
	return db
}
