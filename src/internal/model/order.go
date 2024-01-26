package model

import "time"

const OrderUIDLength = 19

type Order struct {
	OrderUID          string    `json:"order_uid" validate:"required,len=19"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" validate:""`
	CustomerID        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required,numeric"`
	SmID              *int      `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required,numeric"`
}

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required,numeric"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type Payment struct {
	Transaction  string   `json:"transaction" validate:"required"`
	RequestID    string   `json:"request_id" validate:""`
	Currency     string   `json:"currency" validate:"required"`
	Provider     string   `json:"provider" validate:"required"`
	Amount       *int     `json:"amount" validate:"required"`
	PaymentDT    *int64   `json:"payment_dt" validate:"required"`
	Bank         string   `json:"bank" validate:"required"`
	DeliveryCost *float64 `json:"delivery_cost" validate:"required"`
	GoodsTotal   *int     `json:"goods_total" validate:"required"`
	CustomFee    *float64 `json:"custom_fee" validate:"required"`
}

type Item struct {
	ChrtID      *int     `json:"chrt_id" validate:"required"`
	TrackNumber string   `json:"track_number" validate:"required"`
	Price       *float64 `json:"price" validate:"required"`
	RID         string   `json:"rid" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Sale        *int     `json:"sale" validate:"required"`
	Size        string   `json:"size" validate:"required,numeric"`
	TotalPrice  *float64 `json:"total_price" validate:"required"`
	NmID        *int     `json:"nm_id" validate:"required"`
	Brand       string   `json:"brand" validate:"required"`
	Status      *int     `json:"status" validate:"required"`
}
