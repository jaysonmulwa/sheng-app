// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewWord struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	UserID  string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Word struct {
	ID      string `json:"id" bson:"_id"`
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	Author  *User  `json:"author"`
}
