package model

import "time"

type Activity struct {
	UserId int
	Name   string
	Date   time.Time
}
