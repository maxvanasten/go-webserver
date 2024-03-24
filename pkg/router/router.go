package router

import (
    "net/http"
    "html/template"
    "github.com/maxvanasten/go-webserver/pkg/notes"
)

func GetRouter() *http.ServeMux {
    router := http.NewServeMux()
    
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("public/index.html"))
        tmpl.Execute(w, nil)
    })

    // Method based routing with parameters
    router.HandleFunc("GET /api/{command}", CommandHandler)

    notes_db := []notes.Note{}
    router.HandleFunc("GET /notes", notes.GetNoteHandler)
    router.HandleFunc("POST /notes", notes.CreateNoteHandler)

    return router
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
    command := r.PathValue("command")

    w.Write([]byte("Command: " + command))
}
