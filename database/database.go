package database

import (
	"github.com/RodrigoMattosoSilveira/cc_seeder/seeders/person"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dbName := "/private/var/ContasCorrentes/sqlite_dev.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate the Person model
    db.AutoMigrate(person.Person{})

	return db
}
