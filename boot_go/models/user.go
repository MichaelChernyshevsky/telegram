package models

type User struct {
	Username string
	Password string
	Name     string
	Company  string
	Projects []string
}
