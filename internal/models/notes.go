package main

type Notes struct {
	note_id	string	`json: 'note_id'`
	Title	string	`json:'title'`
	Description	string `json:'description'`
	Body		string	`json:'body'`
}