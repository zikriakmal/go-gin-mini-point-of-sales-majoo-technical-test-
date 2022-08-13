package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primary_key,type:bigint(20)"`
	UserName  string    `json:"user_name" gorm:"type:varchar(45) NOT NULL" faker:"last_name,keep"`
	Password  string    `json:"password" gorm:"type:varchar(255) NOT NULL" faker:"password,keep"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	CreatedBy int64     `json:"created_by" gorm:"type:bigint(20) NOT NULL"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedBy int64     `json:"updated_by" gorm:"type:bigint(20) NOT NULL"`
}
