package betting

import (
	"log"
	"testing"
)

func TestRequestListCompetitions(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	list, err := bet.ListCompetitions(Filter{
		Locale:       "en",
		MarketFilter: MarketFilter{
		//InPlayOnly: false,
		},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Println(list)
	}
}
