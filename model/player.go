package model

import "time"

type Player struct {
	ID          int64     `json:"id"`
	Name   int64     `json:"name"`
	Score     int64     `json:"score"`
	Created   time.Time `json:"created_at"`
	Updated   time.Time `json:"updated_at"`
	Deleted   time.Time `json:"deleted_at"`
}

func (u *Player) TableName() string {
	return "player"
}
