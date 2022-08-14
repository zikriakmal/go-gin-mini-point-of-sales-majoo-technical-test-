package models

import "time"

type Merchant struct {
	ID           int64         `json:"id" gorm:"primary_key,type:bigint(20)"`
	UserId       int64         `json:"user_id" gorm:"type:bigint(20) NOT NULL"`
	MerchantName string        `json:"merchant_name" gorm:"type:varchar(40) NOT NULL"`
	CreatedAt    time.Time     `json:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	CreatedBy    int64         `json:"created_by" gorm:"type:bigint(20) NOT NULL"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedBy    int64         `json:"updated_by" gorm:"type:bigint(20) NOT NULL"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:MerchantId;references:ID"`
}
