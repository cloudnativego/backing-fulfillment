package service

type fulfillmentStatus struct {
	SKU             string `json:"sku"`
	ShipsWithin     int    `json:"ships_within"`
	QuantityInStock int    `json:"qty_in_stock"`
}
