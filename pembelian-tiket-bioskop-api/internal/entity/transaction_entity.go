package entity

const (
	PENDING   = "pending"
	CONFIRMED = "confirmed"
	CANCEL    = "cancelled"
)

type Transaction struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AccountID     string `gorm:"not null"`
	ShowtimeID string `gorm:"not null"`
	SeatID     string `gorm:"not null;unique"`
	Status     string `gorm:"not null"`

	Account  Account  `gorm:"foreignKey:AccountID"`
	Showtime Showtime `gorm:"foreignKey:ShowtimeID"`
	Seat     Seat     `gorm:"foreignKey:SeatID"`
}
