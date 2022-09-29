package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	// user:password@/db_name
	// user     -> localhost
	// password ->
	// db_name  -> fcc_bookstore_db
	d, err := gorm.Open("mysql", "localhost:@/fcc_bookstore_db")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
