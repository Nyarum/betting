package betting

import (
	"crypto/rand"
	"crypto/tls"
	"log"

	"github.com/go-resty/resty"
)

const (
	CertURL    = "https://identitysso.betfair.com/api/certlogin"
	AccountURL = "https://api-au.betfair.com/exchange/account/rest/v1.0/"
	BettingURL = "https://api-au.betfair.com/exchange/betting/rest/v1.0/"
)

type BettingFactory interface {
	LoadKeys(string, string) error
	GetSession(string, string) error
	GetAppKeys() ([]DeveloperApp, error)
}

type Betting struct {
	ApiKey     string
	SessionKey string
}

func NewBet(apiKey string) BettingFactory {
	return &Betting{
		ApiKey: apiKey,
	}
}

func (b *Betting) LoadKeys(pemCert, keyCert string) error {
	cert, err := tls.LoadX509KeyPair(pemCert, keyCert)
	if err != nil {
		return err
	}

	ssl := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	ssl.Rand = rand.Reader

	resty.SetTLSClientConfig(ssl)

	return nil
}

type Session struct {
	SessionToken string
	LoginStatus  string
}

func (b *Betting) GetSession(login, password string) error {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"username": login,
		"password": password,
	}).SetHeader("X-Application", b.ApiKey).SetHeader("Content-Type", "application/x-www-form-urlencoded").Post(CertURL)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(resp)

	return nil
}
