package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	// user:password@/db_name?
	// user     -> root
	// password ->
	// db_name  -> fcc_bookstore_db
	user, pwd, dbName := "root", "", "fcc_bookstore_db"
	dbOptions := "charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", user+":"+pwd+"@/"+dbName+"?"+dbOptions)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
