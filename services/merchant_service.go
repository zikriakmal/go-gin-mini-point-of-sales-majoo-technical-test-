package services

import (
	"majoo/dto"
	"majoo/repositories"
)

type MerchantService interface {
	GetAll(userId int64, request dto.MerchantsReqDto) (dto.PaginationRes[dto.MerchantResDto], error)
	Get(userId int64, request dto.MerchantReqDto) (dto.MerchantResDto, error)
}

type merchantService struct {
	merchantRepo repositories.MerchantRepository
}

func NewMerchantService(merchantRepo repositories.MerchantRepository) MerchantService {
	return &merchantService{
		merchantRepo: merchantRepo,
	}
}

func (s *merchantService) GetAll(userId int64, request dto.MerchantsReqDto) (dto.PaginationRes[dto.MerchantResDto], error) {
	var result dto.PaginationRes[dto.MerchantResDto]
	var modelToDto []dto.MerchantResDto
	merchants, total, err := s.merchantRepo.GetAllByMerchantId(userId, request)

	if err != nil {
		return result, err
	}
	for _, v := range merchants {
		omzet := float64(0)
		for _, k := range v.Transactions {
			omzet += k.BillTotal
		}
		modelToDto = append(modelToDto, dto.MerchantResDto{
			ID:           v.ID,
			MerchantName: v.MerchantName,
			Omzet:        omzet,
		})
	}

	result = dto.PaginationRes[dto.MerchantResDto]{
		MetaData: struct {
			Page    int   `json:"page"`
			PerPage int   `json:"per_page"`
			Total   int64 `json:"total"`
		}{
			Page:    request.Page,
			PerPage: request.Limit,
			Total:   total,
		},
		Records: modelToDto,
	}

	return result, nil
}

func (s *merchantService) Get(userId int64, request dto.MerchantReqDto) (dto.MerchantResDto, error) {
	var result dto.MerchantResDto

	merchant, err := s.merchantRepo.GetByMerchantId(userId, request)

	if err != nil {
		return result, err
	}

	omzet := float64(0)
	for _, k := range merchant.Transactions {
		omzet += k.BillTotal
	}

	result = dto.MerchantResDto{
		ID:           merchant.ID,
		MerchantName: merchant.MerchantName,
		Omzet:        omzet,
	}
	return result, nil
}
