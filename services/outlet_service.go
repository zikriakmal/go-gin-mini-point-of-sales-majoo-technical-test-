package services

import (
	"majoo/dto"
	"majoo/repositories"
)

type OutletService interface {
	GetAll(request dto.OutletsReqDto) (dto.PaginationRes[dto.OutletResDto], error)
	Get(request dto.OutletReqDto) (dto.OutletResDto, error)
}

type outletService struct {
	outletRepo repositories.OutletRepository
}

func NewOutletService(outletRepo repositories.OutletRepository) OutletService {
	return &outletService{
		outletRepo: outletRepo,
	}
}

func (s *outletService) GetAll(request dto.OutletsReqDto) (dto.PaginationRes[dto.OutletResDto], error) {
	var result dto.PaginationRes[dto.OutletResDto]
	var modelToDto []dto.OutletResDto

	outlets, total, err := s.outletRepo.GetAllByMerchantId(request)
	if err != nil {
		return result, err
	}

	for _, v := range outlets {
		omzet := float64(0)
		for _, k := range v.Transactions {
			omzet += k.BillTotal
		}
		modelToDto = append(modelToDto, dto.OutletResDto{
			ID:           v.ID,
			MerchantName: v.Merchant.MerchantName,
			OutletName:   v.OutletName,
			Omzet:        omzet,
		})
	}

	result = dto.PaginationRes[dto.OutletResDto]{
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

	return result, err
}

func (s *outletService) Get(request dto.OutletReqDto) (dto.OutletResDto, error) {
	var result dto.OutletResDto

	outlet, err := s.outletRepo.GetByOutletId(request)
	if err != nil {
		return result, err
	}

	omzet := float64(0)
	for _, k := range outlet.Transactions {
		omzet += k.BillTotal
	}

	result = dto.OutletResDto{
		ID:           outlet.ID,
		MerchantName: outlet.Merchant.MerchantName,
		OutletName:   outlet.OutletName,
		Omzet:        omzet,
	}
	return result, err
}
