package api

import (
	"fmt"
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"log"
	"net/http"
	"strings"
)

func HandleGetRoomMembers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		path := r.URL.Path
		parts := strings.Split(path, "/")
		roomId := parts[3]
		log.Print(roomId)

		db, dbErr := utils.ConnectDB()

		if dbErr != nil {
			log.Fatal("Failed to connect to the database: ", dbErr)
		}

		if utils.DoesRoomExist(db, roomId) {
			var room models.Room

			type RoomMembers struct {
				Member_name       string
				Member_student_id string
			}

			// Find the first room with the provided room_pin
			db.First(&room, "room_pin = ?", roomId)

			log.Print(room.RoomMembers)

			fmt.Fprint(w, room.RoomMembers)
		} else {
			http.Error(w, "Room does not exist", http.StatusNotFound)
		}

		return
	}
}
