package betting

import (
	"log"
	"testing"
)

func TestRequestPlaceOrders(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.PlaceOrders(Filter{
		PlaceOrdersMarketID: "1.139017587",
		PlaceOrdersInstructions: []PlaceInstruction{
			PlaceInstruction{
				SelectionID: 256171,
				Side:        S_LAY,
				Handicap:    0,
				OrderType:   OT_MARKET_ON_CLOSE,
				MarketOnCloseOrder: MarketOnCloseOrder{
					Liability: 1,
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestRequestPlaceOrders\n %v \n\n", list)
	}
}
