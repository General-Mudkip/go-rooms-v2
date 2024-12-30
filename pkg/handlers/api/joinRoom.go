package api

import (
	"encoding/json"
	"fmt"
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"log"
	"net/http"
	"strings"
)

func HandleJoinRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db, dbErr := utils.ConnectDB()

		if dbErr != nil {
			log.Fatal("Failed to connect to the database: ", dbErr)
		}

		parseErr := r.ParseForm()
		if parseErr != nil {
			log.Fatal("Failed to parse the form.", parseErr)
		}

		if utils.DoesRoomExist(db, r.FormValue("room_pin")) {
			log.Print("JOINING...")
			// Various needed types
			var room models.Room

			type RoomMembers struct {
				Member_name       string
				Member_student_id string
			}
			var currentMembers []RoomMembers

			// Find the first room with the provided room_pin
			db.First(&room, "room_pin = ?", strings.TrimSpace(r.FormValue("room_pin")))

			// Convert the JSON string into an object we can use
			log.Print("UNMARSHALLING...")
			if room.RoomMembers != "" {
				unmarshalErr := json.Unmarshal([]byte(room.RoomMembers), &currentMembers)

				if unmarshalErr != nil {
					log.Print("ERROR!" + unmarshalErr.Error())
					return
				}
			}

			// Append the new member to the currentMembers list
			currentMembers = append(currentMembers, RoomMembers{
				Member_name:       r.FormValue("student_name"),
				Member_student_id: r.FormValue("student_id"),
			})

			log.Print(currentMembers)

			// Convert the object back to a JSON string
			data, _ := json.Marshal(currentMembers)
			room.RoomMembers = string(data)

			log.Print("SAVING...")
			// Update it
			db.Save(&room)

			fmt.Fprint(w, "<span class='text-green-400'>success</span>")
		} else {
			fmt.Fprint(w, "<span class='text-red-400'>Room does not exist.</span>")
		}

		return
	}
}
