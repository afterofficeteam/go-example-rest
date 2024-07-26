package model

import "github.com/google/uuid"

type Book struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	URL     string    `json:"url"`
	Authors []Author  `json:"authors"`
}

type Author struct {
	Name string `json:"name`
	URL  string `json:"url`
}
