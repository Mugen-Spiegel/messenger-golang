package models

type Contact struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ContactID string `json:"contact_id" gorm:"primaryKey"`
	Contact   User
	UserID    string `json:"user_id" gorm:"primaryKey"`
	User      User
}
