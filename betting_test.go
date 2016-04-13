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
	Debug    bool   `json:"debug"`
}

func loadConfig() (test Test) {
	loadFile, err := ioutil.ReadFile("testdata/test.json")
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

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	err = bet.KeepAlive()
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

	keys, err := bet.GetAppKeys()
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Println(keys)
	}
}

func TestRequestAccountDetails(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	details, err := bet.GetAccountDetails()
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Println(details)
	}
}

func TestRequestAccountFunds(t *testing.T) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		t.Error(err)
	}

	funds, err := bet.GetAccountFunds(Filter{Wallet: UK})
	if err != nil {
		t.Error(err)
	}

	if config.Debug {
		log.Println(funds)
	}
}

func BenchmarkRequestKeys(b *testing.B) {
	config := loadConfig()

	bet := NewBet(config.ApiKey)

	err := bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		_, err = bet.GetAppKeys()
		if err != nil {
			b.Error(err)
		}
	}
}
