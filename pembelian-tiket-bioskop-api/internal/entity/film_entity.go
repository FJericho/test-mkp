package entity

type Film struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string `gorm:"not null"`
	Genre       string `gorm:"not null"`
	Duration    int    `gorm:"not null"`
	Description string `gorm:"type:text"`

	Showtime  []Showtime `gorm:"foreignKey:FilmID"`
}
