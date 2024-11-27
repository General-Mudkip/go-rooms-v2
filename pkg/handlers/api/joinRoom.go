package api

import (
	"fmt"
	"go-rooms/pkg/utils"
	"log"
	"net/http"
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

		log.Print("Form value pin: " + r.FormValue("room_pin"))

		if utils.DoesRoomExist(db, r.FormValue("room_pin")) {
			fmt.Fprint(w, "<span class='text-green-400'>does exist</span>")
		} else {
			fmt.Fprint(w, "<span class='text-red-400'>no exist</span>")
		}

		return
	}
}
