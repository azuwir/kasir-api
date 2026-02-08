package models

type Transaction struct {
	ID          int                  `json:"id"`
	TotalAmount float64              `json:"total_amount"`
	Details     []TransactionDetails `json:"details"`
}

type TransactionDetails struct {
	ID            int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name,omitempty"`
	Quantity      int     `json:"quantity"`
	Subtotal      float64 `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}
