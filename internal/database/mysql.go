package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ridhopujiono/nusanfood-api/internal/config"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Get("DB_USER", "root"),
		config.Get("DB_PASS", ""),
		config.Get("DB_HOST", "localhost"),
		config.Get("DB_PORT", "3306"),
		config.Get("DB_NAME", "db_recipe"),
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB = db
}
