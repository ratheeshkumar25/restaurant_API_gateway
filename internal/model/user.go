package model

type UserModel struct {
	UserID int    `gorm:"primaryKey;autoIncrement"`
	Phone  string `json:"phone" validate:"require"`
}

type VerifyOTP struct {
	Phone string `json:"phone"`
	Otp   string `json:"otp"`
}
