package domain

type Currency struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Nama string `gorm:"type:varchar(15);not null"`
}
