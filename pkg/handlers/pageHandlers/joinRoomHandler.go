package pageHandlers

import (
	"net/http"
	"text/template"
)

func HandleJoinRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./pkg/templates/pages/joinRoom.html"))
		tmpl.Execute(w, map[string]string{"Title": "HTMX with Go"})
	}
}
