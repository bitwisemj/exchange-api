package service

import (
	"exchange-api/pkg/external/dto"
	"log"
)

func GetQuotationOrchestratorService(currency string) (*dto.QuotationBidDTO, error) {

	quotationDTO, err := GetDollarQuotation(currency)

	if err != nil {
		log.Fatalf("Could not get quotation, error: %v", err)
		return nil, err
	}

	err = CreateQuotation(quotationDTO)

	if err != nil {
		log.Fatalf("Could not create quotation, error: %v", err)
		return nil, err
	}

	quotationBidDTO := dto.QuotationBidDTO{
		Bid: quotationDTO.USDBRL.Bid,
	}

	return &quotationBidDTO, nil
}
