package entity

const (
	AVAILABLE = "available"
	BOOKED    = "booked"
	CANCELLED = "cancelled"
)

type Seat struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ShowtimeID string `gorm:"not null"`
	SeatNumber string `gorm:"not null"`
	Status     string `gorm:"not null"`

	Showtime Showtime `gorm:"foreignKey:ShowtimeID"`
}
