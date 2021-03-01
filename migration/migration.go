package migration

import (
	"github.com/davidyap2002/user-go/config"
	"github.com/davidyap2002/user-go/graph/model"
)

//MigrateTable Migrate Table
func MigrateTable() {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User
	var fileUpload model.FileUpload
	var igPost model.IgPost
	var igPostFile model.IgPostFile

	db.AutoMigrate(&user)
	db.AutoMigrate(&fileUpload)
	db.AutoMigrate(&igPost)
	db.AutoMigrate(&igPostFile)
}
