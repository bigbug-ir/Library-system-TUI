package models

type Book struct {
	ID        int
	Title     string
	Author    string
	ISBN      string
	Copies    int
	Available bool
}
