package entities

import "time"

type Entity struct {
	ID        int64     `json:"-"`
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}
