package models

import "time"

type ModelPhoto struct {
	UserId    string    `bson:"_id"`
	NameImg   string    `json:nameImg`
	Size      int64     `json:size`
	CreatedAt time.Time `json:createdAt`
}
