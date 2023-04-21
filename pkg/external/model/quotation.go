package model

import "gorm.io/gorm"

type Quotation struct {
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	CreateDate string
	gorm.Model
}
