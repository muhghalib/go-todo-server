package entities

import (
	"time"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
	Overdue    TaskStatus = "overdue"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"size:100;not null"`
	Description string     `json:"description" gorm:"type:text;not null"`
	Status      TaskStatus `json:"status" gorm:"type:enum('pending','in_progress','completed','overdue');default:'pending'"`
	DueDate     time.Time  `json:"dueDate" gorm:"index"`
	UserID      uint       `json:"userId"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}
