package models

import "time"

type Transaction struct {
	ID         int64     `json:"id" gorm:"primary_key,type:bigint(20)"`
	MerchantId int64     `json:"merchant_id" gorm:"type:bigint(20) NOT NULL"`
	OutletId   int64     `json:"outlet_id" gorm:"type:bigint(20) NOT NULL"`
	BillTotal  float64   `json:"bill_total" gorm:"type:double NOT NULL"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	CreatedBy  int64     `json:"created_by" gorm:"type:bigint(20) NOT NULL"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedBy  int64     `json:"updated_by" gorm:"type:bigint(20) NOT NULL"`
	Merchant   *Merchant `json:"merchant" gorm:"foreignKey:MerchantId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Outlet     *Outlet   `json:"outlet" gorm:"foreignKey:OutletId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
