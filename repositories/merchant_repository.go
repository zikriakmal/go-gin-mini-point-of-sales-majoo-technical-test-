package repositories

import (
	"gorm.io/gorm"
	"majoo/dto"
	"majoo/helper"
	"majoo/models"
)

type MerchantRepository interface {
	GetAllByMerchantId(userId int64, request dto.MerchantsReqDto) ([]models.Merchant, int64, error)
	GetByMerchantId(userId int64, request dto.MerchantReqDto) (models.Merchant, error)
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{
		db: db,
	}
}

type merchantRepository struct {
	db *gorm.DB
}

func (o merchantRepository) GetAllByMerchantId(userId int64, request dto.MerchantsReqDto) ([]models.Merchant, int64, error) {
	var merchants []models.Merchant
	var count int64

	o.db.Model(&models.Merchant{}).Count(&count).
		Where("user_id = ? ", userId)

	dateFrom := helper.SetDefaultDate(request.DateFrom, 2022, 8, 1)
	dateTo := helper.SetDefaultDate(request.DateFrom, 2022, 10, 30)

	err := o.db.Model(&models.Merchant{}).Scopes(helper.Paginate(request.PaginationReqDto)).
		Preload("Transactions", func(db *gorm.DB) *gorm.DB {
			return db.Where("created_at >= ?", dateFrom).
				Where("created_at <= ?", dateTo)
		}).
		Where("user_id = ?", userId).
		Find(&merchants).Error

	return merchants, count, err
}

func (o merchantRepository) GetByMerchantId(userId int64, request dto.MerchantReqDto) (models.Merchant, error) {
	var merchant models.Merchant
	err := o.db.Where("id = ? AND user_id = ?", request.MerchantId, userId).First(&merchant).Error
	return merchant, err
}
