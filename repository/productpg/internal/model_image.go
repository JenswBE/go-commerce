package internal

type Image struct {
	ID        string `gorm:"type:uuid"`
	OwnerID   string `gorm:"type:uuid"`
	OwnerType string

	Extension string // File extension
	Order     int
}
