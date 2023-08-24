package internal

type Event struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}
