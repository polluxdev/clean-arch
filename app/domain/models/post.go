package models

type Post struct {
	ID     int8   `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}
