package services

import (
	"majoo/dto"
	"majoo/repositories"
	"time"
)

type MerchantService interface {
	GetAll(userId int64, request dto.MerchantsReqDto) (dto.PaginationRes[dto.MerchantResDto], error)
	Get(userId int64, request dto.MerchantReqDto) (dto.MerchantResDto, error)
	GetReport(userId int64, request dto.MerchantReportReqDto) (dto.PaginationRes[dto.MerchantReportResDto], error)
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

func (s *merchantService) GetReport(userId int64, request dto.MerchantReportReqDto) (dto.PaginationRes[dto.MerchantReportResDto], error) {
	var result dto.PaginationRes[dto.MerchantReportResDto]

	merchant, err := s.merchantRepo.GetReportsByMerchantId(userId, request)

	if err != nil {
		return result, err
	}

	var finalRes []dto.MerchantReportResDto

	total := int64((request.DateTo.Sub(request.DateFrom).Hours() / 24) + 1)
	for i := int64(request.DateFrom.Day()); i <= int64(request.DateFrom.Day())+total-1; i++ {
		var sumDaily float64 = 0
		for _, v := range merchant.Transactions {
			if v.CreatedAt.Month() == request.DateFrom.Month() && v.CreatedAt.Day() == int(i) {
				sumDaily += v.BillTotal
			}
		}
		finalRes = append(finalRes, dto.MerchantReportResDto{
			MerchantName: merchant.MerchantName,
			Omzet:        sumDaily,
			Date:         time.Date(request.DateTo.Year(), request.DateTo.Month(), int(i), 0, 0, 0, 0, time.UTC),
		})
	}

	start := (request.Page * request.Limit) - request.Limit
	end := request.Limit * request.Page

	if start <= int(total) && end <= int(total) {
		finalRes = finalRes[start:end]
	} else if start <= int(total) && end >= int(total) {
		finalRes = finalRes
	} else {
		finalRes = []dto.MerchantReportResDto{}
	}

	result = dto.PaginationRes[dto.MerchantReportResDto]{
		MetaData: struct {
			Page    int   `json:"page"`
			PerPage int   `json:"per_page"`
			Total   int64 `json:"total"`
		}{
			Page:    request.Page,
			PerPage: request.Limit,
			Total:   total,
		},
		Records: finalRes,
	}

	return result, nil
}
