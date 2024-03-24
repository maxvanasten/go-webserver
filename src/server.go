package main

import (
	"log"
	"net/http"
)

func main() {
	router := get_router()

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
    log.Println("Server is running on port 3000 (http://localhost:3000)")
    server.ListenAndServe()
}

func get_router() *http.ServeMux {
    router := http.NewServeMux()
    
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Method based routing with parameters
    router.HandleFunc("POST /api/{command}", CommandHandler)

    return router
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
    command := r.PathValue("command")

    w.Write([]byte("Command: " + command))
}
