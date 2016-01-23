package betting

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

type Test struct {
	ApiKey   string `json:"api_key"`
	Login    string `json:"login"`
	Password string `json:"password"`
	CertPem  string `json:"cert_pem"`
	CertKey  string `json:"cert_key"`
}

func loadConfig() (test Test) {
	loadFile, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(loadFile, &test)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func TestRequestLogin(t *testing.T) {
	config := loadConfig()

	err := NewBet(config.ApiKey).GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}
}

func TestRequestAppKeys(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	_, err = bet.GetAppKeys()
	if err != nil {
		t.Error(err)
	}
}
