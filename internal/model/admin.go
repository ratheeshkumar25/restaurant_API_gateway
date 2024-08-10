package model

// AdminModel represents the admin entity.
type AdminModel struct {
	AdminID  int    `gorm:"primaryKey" `
	Username string `json:"username"`
	Password string `json:"password"`
}
