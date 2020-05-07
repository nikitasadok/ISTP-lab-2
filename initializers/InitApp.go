package initializers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/validations"
)

var Db *gorm.DB

func InitApp(connectionString string) (*gorm.DB, error) {
	var err error
	Db, err = gorm.Open("mysql", connectionString)
	if Db == nil {
		print("Err opening")
	}

	validations.RegisterCallbacks(Db)
	return Db, err
}
