package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Email       string         `gorm:"column:email;type:varchar(255);uniqueIndex"`
	Password    string         `gorm:"column:password;type:varchar(255)"`
	LastLoginAt *time.Time     `gorm:"column:last_login_at;type:timestamp"`
	LastLoginIP string         `gorm:"column:last_login_ip;type:varchar(255)"`
	Permission  uint           `gorm:"column:permission;type:int;default:0"`
	IsBanned    *bool          `gorm:"column:is_banned;type:boolean;default:false"`
	CreatedAt   time.Time      `gorm:"created_at:email;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"updated_at:email;autoUpdateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"delete_at:email;index"`
}
