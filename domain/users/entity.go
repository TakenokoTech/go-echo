package users

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UserName  string    `gorm:"column:username" json:"UserName" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
