package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = []User{
	{ID: "1", Username: "dj", Password: "SamplePassword123!"},
	{ID: "2", Username: "western", Password: "ThisIsMyPassword"},
	{ID: "3", Username: "akunamatata", Password: "LyonKing1!"},
}
