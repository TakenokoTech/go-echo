package todos

import "time"

type Todo struct {
	ID        uint      `gorm:"column:id primaryKey"`
	UserId    string    `gorm:"column:user_id" json:"UserId" validate:"required"`
	Task      string    `gorm:"column:task" json:"Task" validate:"required"`
	Status    string    `gorm:"column:status" json:"Status" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
