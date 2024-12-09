package entities

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"size:100;unique;not null"`
	Password  string    `json:"password" gorm:"size:100;not null"`
	Role      string    `json:"role" gorm:"type:enum('default','role');default:'default'"`
	Tasks     *[]Task   `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
