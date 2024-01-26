package model

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderExists   = errors.New("order with this uid already exists")
	ErrInvalidData   = errors.New("received data is not valid")
)
