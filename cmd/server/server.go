package server

import (
	"exchange-api/pkg/external/api"
	"net/http"
)

func StartServer() {

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", api.GetQuotationController)
	http.ListenAndServe(":8080", mux)
}
