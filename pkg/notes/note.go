package notes

import (
	"encoding/json"
	"net/http"
    "html/template"
)

var notes_db = []Note{}

// Note is a struct that represents a note
type Note struct {
	Title string
	Body  string
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    
    notes_to_send := map[string][]Note{
        "Notes": {
            {"Note 1", "This is the body of note 1"},
            {"Note 2", "This is the body of note 2"},
            {"Note 3", "This is the body of note 3"},
        },
    }

	tmpl := template.Must(template.ParseFiles("public/notes.html"))
	tmpl.Execute(w, notes_to_send)
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
