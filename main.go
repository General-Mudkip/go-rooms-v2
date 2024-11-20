package main

import (
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/base.html", "templates/pages/index.html")
	tmpl.Execute(w, map[string]string{"Title": "HTMX with Go"})
}

func main() {
	db, err := utils.ConnectDB()
	log.Print("HELLO!")

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	crErr := utils.CreateRoomsTable(db)
	if crErr != nil {
		log.Fatal("Failed to create rooms table: ", err)
	}

	var rooms []models.Room
	db.Create(&models.Room{RoomName: "Test Palace", RoomPin: "1234", RoomMembers: ""})
	db.Find(&rooms)
	log.Print(rooms)

	log.Print("After create")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
