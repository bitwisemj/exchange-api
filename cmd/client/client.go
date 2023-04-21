package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type QuotationResponseDTO struct {
	Bid string `json:"bid"`
}

func main() {

	quotationDTO, err := GetDollarQuotation()

	if err != nil {
		panic(err)
	}

	SaveQuotation(*quotationDTO)
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(quotationDTO)
}

func GetDollarQuotation() (*QuotationResponseDTO, error) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, (5000 * time.Millisecond))

	defer cancel()

	client := http.Client{}
	request, err := http.
		NewRequestWithContext(
			ctx,
			"GET",
			"http://localhost:8080/cotacao?currency=USD-BRL",
			nil,
		)

	if err != nil {
		panic(err)
	}

	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("Could not execute http request, error: %v", err)
		return nil, err
	}

	defer response.Body.Close()

	byteArr, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Could not parse response body, error: %v", err)
		return nil, err
	}

	quotationDTO := QuotationResponseDTO{}
	json.Unmarshal(byteArr, &quotationDTO)
	return &quotationDTO, nil
}

func SaveQuotation(quotationDTO QuotationResponseDTO) {

	file, err := os.Create("cotacao.txt")

	if err != nil {
		log.Fatalf("Could not open file, error: %v", err)
		panic(err)
	}

	defer file.Close()

	content := fmt.Sprintf("Dólar: %v", quotationDTO.Bid)
	file.WriteString(content)
}
