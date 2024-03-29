package repositories

import (
	"gorm.io/gorm"
	"majoo/dto"
	"majoo/helper"
	"majoo/models"
)

type OutletRepository interface {
	GetAllByMerchantId(request dto.OutletsReqDto) ([]models.Outlet, int64, error)
	GetByOutletId(request dto.OutletReqDto) (models.Outlet, error)
}

func NewOutletRepository(db *gorm.DB) OutletRepository {
	return &outletRepository{
		db: db,
	}
}

type outletRepository struct {
	db *gorm.DB
}

func (o outletRepository) GetAllByMerchantId(request dto.OutletsReqDto) ([]models.Outlet, int64, error) {
	var outlets []models.Outlet
	var count int64

	o.db.Model(&models.Outlet{}).Count(&count).
		Where("merchant_id = ?", request.MerchantId)

	dateFrom := helper.SetDefaultDate(request.DateFrom, 2022, 8, 1)
	dateTo := helper.SetDefaultDate(request.DateFrom, 2022, 10, 30)

	err := o.db.Model(&models.Outlet{}).Scopes(helper.Paginate(request.PaginationReqDto)).
		Preload("Merchant").
		Preload("Transactions", func(db *gorm.DB) *gorm.DB {
			return db.Where("created_at >= ?", dateFrom).
				Where("created_at <= ?", dateTo)
		}).
		Where("merchant_id = ?", request.MerchantId).
		Find(&outlets).Error

	return outlets, count, err
}

func (o outletRepository) GetByOutletId(request dto.OutletReqDto) (models.Outlet, error) {
	var outlet models.Outlet
	err := o.db.
		Preload("Merchant").
		Preload("Transactions").
		Where("merchant_id = ? AND id = ? ", request.MerchantId, request.OutletId).
		First(&outlet).
		Error
	return outlet, err
}
