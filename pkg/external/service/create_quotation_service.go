package service

import (
	"context"
	"exchange-api/pkg/config"
	"exchange-api/pkg/external/dto"
	"exchange-api/pkg/external/model"
	"log"
	"time"
)

func CreateQuotation(quotationDTO *dto.QuotationResponseDTO) error {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, (time.Millisecond * 100))
	defer cancel()
	db, err := config.GetConnection()

	if err != nil {
		log.Fatalf("Could not establish database connection, error: %v", err)
		return err
	}

	quotation := model.Quotation{
		Code:       quotationDTO.USDBRL.Code,
		Codein:     quotationDTO.USDBRL.Codein,
		Name:       quotationDTO.USDBRL.Name,
		High:       quotationDTO.USDBRL.High,
		Low:        quotationDTO.USDBRL.Low,
		VarBid:     quotationDTO.USDBRL.VarBid,
		PctChange:  quotationDTO.USDBRL.PctChange,
		Bid:        quotationDTO.USDBRL.Bid,
		Ask:        quotationDTO.USDBRL.Ask,
		CreateDate: quotationDTO.USDBRL.CreateDate,
	}

	db.WithContext(ctx).Create(&quotation)
	log.Printf("Quotation created successfully with id %v", quotation.ID)
	return nil
}
