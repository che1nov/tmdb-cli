package models

import "time"

type Movie struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	TMDBID      int  `gorm:"uniqueIndex"`
	Title       string
	Overview    string
	ReleaseDate string
	VoteAverage float32
	CreatedAt   time.Time
}
