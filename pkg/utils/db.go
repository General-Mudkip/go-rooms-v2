package utils

import (
	"go-rooms/pkg/models"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Sets up a connection to the database. Returns the database object.
func ConnectDB() (*gorm.DB, error) {
	dbURL := "./dev.db"
	// dbToken := os.Getenv("DATABASE_TOKEN")

	// db, err := gorm.Open(sqlite.Open(dbURL+"?authToken="+dbToken), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateRoomsTable(db *gorm.DB) error {
	log.Print("in create")
	err := db.AutoMigrate(&models.Room{})

	return err
}
