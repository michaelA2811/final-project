package entity

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null;unique;type:varchar(191)"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Password  string `gorm:"not null;unique;type:varchar(191)"`
	Age       int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Token struct {
// 	Token string `json:"token"`
// }

// type Credentials struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type Claims struct {
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }
