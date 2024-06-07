package config 


import {
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
}


var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:root/simplerest?charset-utf8&parseTime=True&loc=Local")
	if err!= nil {
		panic("Failed to connect to database! %s", err)
	}
	db = d
}

fucn GetDB() *gorm.db(
	return db
)