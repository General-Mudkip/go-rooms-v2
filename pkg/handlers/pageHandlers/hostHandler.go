package pageHandlers

import (
	"fmt"
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func HandleHostRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := r.URL.Path
		parts := strings.Split(path, "/")

		// Expect the path to be /host-room/{roomID}
		if len(parts) < 3 || parts[2] == "" {
			http.Error(w, "Room ID not provided", http.StatusBadRequest)
			return
		}

		tmpl := template.Must(template.ParseFiles("./pkg/templates/pages/hostRoom.html"))
		tmpl.Execute(w, map[string]string{"Title": "HTMX with Go"})
	} else {
		db, err := utils.ConnectDB()

		if err != nil {
			log.Fatal("Failed to connect to the database: ", err)
		}

		var rooms []models.Room
		db.Find(&rooms)
		fmt.Fprint(w, rooms)
		return
	}
}
