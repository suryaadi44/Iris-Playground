package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()"`
	Email       string         `gorm:"type:varchar(255);uniqueIndex"`
	Username    string         `gorm:"type:varchar(255);uniqueIndex"`
	Password    string         `gorm:"type:varchar(255)"`
	LastLoginAt *time.Time     `gorm:"type:timestamp"`
	LastLoginIP string         `gorm:"type:varchar(255)"`
	Permission  uint           `gorm:"type:int;default:0"`
	IsBanned    *bool          `gorm:"type:boolean;default:false"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
}
