package model

type Lead struct {
	ID        int    `gorm:"column:id"`
	FirstName string `gorm:"column:firstname"`
	LastName  string `gorm:"column:lastname"`
	Text      string `gorm:"column:fabule"`
	Phone     string `gorm:"column:intphone"`
	City      string `gorm:"column:city"`
}
