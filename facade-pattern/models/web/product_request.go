package web

import "go-design-pattern/facade-pattern/models/domain"

type ProductToBuyRequest struct {
	Product  domain.Product `json:"product"`
	Quantity uint           `json:"quantity"`
}
