package api

import (
	"encoding/json"
	"exchange-api/pkg/external/service"
	"net/http"
)

func GetQuotationController(response http.ResponseWriter, request *http.Request) {

	currency := request.URL.Query().Get("currency")
	quotationDTO, err := service.GetQuotationOrchestratorService(currency)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(response)
	encoder.Encode(quotationDTO)
}
