// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InformNonBeliverInput struct {
	Phone string `json:"phone"`
}

type InformNonBeliverPayload struct {
	User    *User  `json:"user"`
	Message string `json:"message"`
}

type User struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
}
