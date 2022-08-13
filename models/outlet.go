package models

import "time"

type Outlet struct {
	ID         int64     `json:"id" gorm:"primary_key,type:bigint(20)"`
	MerchantId int64     `json:"merchant_id" gorm:"type:bigint(20) NOT NULL"`
	OutletName string    `json:"outlet_name" gorm:"type:double NOT NULL"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	CreatedBy  int64     `json:"created_by" gorm:"type:bigint(20) NOT NULL"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedBy  int64     `json:"updated_by" gorm:"type:bigint(20) NOT NULL"`
	Merchant   *Merchant `json:"merchant" gorm:"foreignKey:MerchantId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
