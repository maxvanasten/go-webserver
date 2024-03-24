package notes

import (
    "time"
    "net/http"
)

var notes_db = []Note{}

// Note is a struct that represents a note
type Note struct {
    Title string
    Body string
    CreatedAt time.Time
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("GET /notes"))
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("POST /notes"))
}
