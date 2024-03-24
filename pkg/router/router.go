package router

import (
    "net/http"
    "html/template"
)

func GetRouter() *http.ServeMux {
    router := http.NewServeMux()
    
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("public/index.html"))
        tmpl.Execute(w, nil)
    })

    // Method based routing with parameters
    router.HandleFunc("GET /api/{command}", CommandHandler)

    return router
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
    command := r.PathValue("command")

    w.Write([]byte("Command: " + command))
}
