package models


type User struct{
	UserName	string	`bson:"username" json: username`
	Password	string	`bson: "password" json: password`
	Role		string	
}

