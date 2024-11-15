package models

import "time"

type Message struct {
	Id        uint64    `json:"id" db:"id"`
	Sender    string    `json:"sender" db:"sender"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Content   string    `json:"content" db:"content"`
}
