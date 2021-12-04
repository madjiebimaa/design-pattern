package web

import "go-design-pattern/facade-pattern/models/domain"

type ProductToBuyResponse struct {
	Product domain.Product `json:"product"`
	Message string         `json:"message"`
}
