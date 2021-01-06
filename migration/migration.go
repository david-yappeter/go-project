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

	db.AutoMigrate(&user)
}
