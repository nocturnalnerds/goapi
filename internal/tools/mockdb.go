package tools

import (
	"time"
)

type mockDb struct{}

var mockLoginDetails = map[string]loginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jason": {
		Coins: 200,
		Username: "jason",
	},
	"marie": {
		Coins: 300,
		Username: "marie",
	},
}

func (d * mockDb) GetUserLoginDetails(username string) *loginDetails{
	
	time.Sleep(time.Second + 1)

	var clientData = loginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d * mockDb) GetUserCoins(username string) *CoinDetails{
	
	time.Sleep(time.Second + 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDb) SetupDatabase() error{
	return nil;
}