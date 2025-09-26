// User model -> defines the user struct
package main

type User struct {
	user_id		string	`json:'user_id'`
	first_name	string	`json:'first_name'`
	last_name	string	`json:'last_name'`
}