package notes

import (
    "net/http"
    "encoding/json"
)

var notes_db = []Note{}

// Note is a struct that represents a note
type Note struct {
    Title string
    Body string
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("GET /notes"))
    for i := range notes_db {
        w.Write([]byte(notes_db[i].Title))
    }
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
    var note Note
    err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    notes_db = append(notes_db, note)
    w.Write([]byte("POST /notes"))
    w.Write([]byte(note.Title))
    w.Write([]byte(note.Body))
}
