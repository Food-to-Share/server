package migrations

import (
	"github.com/Food-to-Share/server/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
