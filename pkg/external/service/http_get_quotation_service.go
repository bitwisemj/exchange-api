package service

import (
	"context"
	"encoding/json"
	"exchange-api/pkg/external/dto"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetDollarQuotation(currency string) (*dto.QuotationResponseDTO, error) {

	url := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%v", currency)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, (time.Millisecond * 800))
	defer cancel()

	client := http.Client{}
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		log.Fatalf("Could not create request with context, error %v", err)
		return nil, err
	}

	response, err := client.Do(request)
	log.Printf("Request done with status %v", response.StatusCode)

	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatalf("Could not execute http request, error: %v", err)
		return nil, err
	}

	defer response.Body.Close()

	byteArr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Could not execute http request, error: %v", err)
		return nil, err
	}

	quotation := dto.QuotationResponseDTO{}
	err = json.Unmarshal(byteArr, &quotation)

	if err != nil {
		log.Fatalf("Could not parse json response, error: %v", err)
		return nil, err
	}

	return &quotation, nil
}
