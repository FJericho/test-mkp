package entity

import "time"

type Showtime struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FilmID    string    `gorm:"not null"`
	StudioID  string    `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`

	Film   Film   `gorm:"foreignKey:FilmID"`
	Studio Studio `gorm:"foreignKey:StudioID"`

	Transactions []Transaction `gorm:"foreignKey:ShowtimeID"`
	Seats        []Seat        `gorm:"foreignKey:ShowtimeID"`
}
