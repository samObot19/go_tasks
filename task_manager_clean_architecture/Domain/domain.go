package domain


import (
	"time"
)

type Task struct{
	ID		string		`bson:"_id,omitempty" json: "id,omitempty"`
	Title		string		`json: title`
	Description	string		`json: description`
	DueDate		time.Time	`json: duedate`
	Status		string		`json: status`
}

type User struct{
	UserName	string	`bson:"username" json: username`
	Password	string	`bson: "password" json: password`
	Role		string
}
