package entity

type Studio struct {
	ID          string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`

	Showtime []Showtime `gorm:"foreignKey:StudioID"`
}
