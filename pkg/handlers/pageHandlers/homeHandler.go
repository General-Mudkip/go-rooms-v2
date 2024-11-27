package pageHandlers

import (
	"fmt"
	"go-rooms/pkg/models"
	"go-rooms/pkg/utils"
	"log"
	"net/http"
	"text/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./pkg/templates/pages/index.html"))
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
