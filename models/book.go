package models

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   int    `json:"page"`
	Rate   int    `json:"rate"`
}
