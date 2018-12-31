package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// CreateTables sets up DB tables
func CreateTables() {
	if err := gorm.DBCon().AutoMigrate(&models.User{}).Error; err != nil {
		panic(err)
	}

	if err := gorm.DBCon().AutoMigrate(&models.TodoList{}).Error; err != nil {
		panic(err)
	}

	if err := gorm.DBCon().AutoMigrate(&models.Todo{}).Error; err != nil {
		panic(err)
	}
}
