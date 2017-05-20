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
		PlaceOrdersMarketID: "1.111836557",
		PlaceOrdersInstructions: []PlaceInstruction{
			PlaceInstruction{
				SelectionID: 5404312,
				Side:        S_BACK,
				OrderType:   OT_MARKET_ON_CLOSE,
				MarketOnCloseOrder: MarketOnCloseOrder{
					Liability: 2,
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
