package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct {
	// Response Code
	Code int
	
	// Account Balance
	Balance int64
}

type Error struct{
	// Error Code
	Code int

	// Error Message
	Message string
}