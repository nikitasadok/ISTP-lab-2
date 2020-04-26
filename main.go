package main

import (
	"CSGORest/initializers"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"CSGORest/routes"
	"net/http"
)

var Db *gorm.DB

func main() {
	_, err := initializers.InitApp("root:nikita@/csgo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	router := routes.NewRouter()
	http.ListenAndServe(":8080", router)
}
