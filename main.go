package main

import (
	"go-rooms/pkg/handlers/api"
	"go-rooms/pkg/handlers/pageHandlers"
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting...")
	db, err := utils.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	crErr := utils.CreateRoomsTable(db)
	if crErr != nil {
		log.Fatal("Failed to create rooms table: ", err)
	}

	var rooms []models.Room
	db.Find(&rooms)

	http.HandleFunc("/", pageHandlers.HandleHome)
	http.HandleFunc("/join-room", pageHandlers.HandleJoinRoom)

	http.HandleFunc("/api/create-room", api.HandleCreateRoom)
	http.HandleFunc("/api/join-room", api.HandleJoinRoom)

	log.Print("Serving...")
	http.ListenAndServe(":8080", nil)
}
