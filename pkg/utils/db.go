package utils

import (
	"go-rooms/pkg/models"
	"strings"

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
	err := db.AutoMigrate(&models.Room{})

	return err
}

func DoesRoomExist(db *gorm.DB, roomPin string) bool {
	var rooms []models.Room
	db.Find(&rooms, "room_pin = ?", strings.TrimSpace(roomPin))

	if len(rooms) != 0 {
		return true
	} else {
		return false
	}
}
