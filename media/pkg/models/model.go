package models

import "time"

type ModelPhoto struct {
	UserId    string    `json:"userId"`
	NameImg   string    `json:nameImg`
	Size      int64     `json:size`
	CreatedAt time.Time `json:createdAt`
}

type ModelFile struct {
	UserId    string    `json:"userId"`
	NameFile  string    `json:contract`
	Size      int64     `json:size`
	CreatedAt time.Time `json:createdAt`
}
