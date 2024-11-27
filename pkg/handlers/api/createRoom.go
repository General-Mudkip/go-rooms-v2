package api

import (
	"fmt"
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func HandleCreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db, dbErr := utils.ConnectDB()

		if dbErr != nil {
			log.Fatal("Failed to connect to the database: ", dbErr)
		}

		parseErr := r.ParseForm()
		if parseErr != nil {
			log.Fatal("Failed to parse the form.", parseErr)
		}

		randomPin := "000000"

		for doWhile := true; doWhile; doWhile = utils.DoesRoomExist(db, randomPin) {
			randomPinInt := rand.Intn(1000000)
			randomPin = strconv.Itoa(randomPinInt)
		}

		// var rooms []models.Room
		db.Create(&models.Room{RoomName: r.FormValue("room_name"), RoomPin: randomPin, RoomMembers: ""})
		fmt.Fprint(w, "<span class='text-red-400'>Success</span>")
		return
	}
}
