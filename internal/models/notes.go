package models

type Notes struct {
	NoteID      string `json:"note_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}
