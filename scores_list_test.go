package betting

import (
	"testing"
	"log"
)

func TestListAvailableEvents(t *testing.T) {
	config := loadConfig()

	betfair := NewBetfair(config.ApiKey)

	err := betfair.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	result, err := betfair.ListAvailableEvents(ScoresParams{
		EventIDs: []string{"28302966"},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestListAvailableEvents\n %v \n\n", result)
	}
}

func TestListScores(t *testing.T) {
	config := loadConfig()

	betfair := NewBetfair(config.ApiKey)

	err := betfair.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	result, err := betfair.ListScores(ScoresParams{
		UpdateKeys: []UpdateKey{{EventID: "28302966"}},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestListScores\n %v \n\n", result)
	}
}

func TestListIncidents(t *testing.T) {
	config := loadConfig()

	betfair := NewBetfair(config.ApiKey)

	err := betfair.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	result, err := betfair.ListIncidents(ScoresParams{
		UpdateKeys: []UpdateKey{{EventID: "28302966"}},
	})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Printf("TestListIncidents\n %v \n\n", result)
	}
}
