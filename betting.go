package betting

import (
	"bytes"
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

// BetURL type of all betfair urls
type BetURL string

const (
	CertURL             = "https://identitysso-api.betfair.com/api/certlogin"
	KeepAliveURL        = "https://identitysso.betfair.com/api/keepAlive"
	AccountURL   BetURL = "https://api.betfair.com/exchange/account/rest/v1.0"
	BettingURL   BetURL = "https://api.betfair.com/exchange/betting/rest/v1.0"
)

// Betting main struct of all method for working with betfair
type Betting struct {
	ApiKey     string
	SessionKey string
}

// NewBet create pointer to Betting struct with base values
func NewBet(apiKey string) *Betting {
	return &Betting{
		ApiKey: apiKey,
	}
}

// Request function for send requests to betfair via REST JSON
func (b *Betting) Request(reqStruct interface{}, url BetURL, method string, filter *Filter) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	urlBuild := bytes.NewBuffer([]byte{})
	urlBuild.WriteString(string(url))
	urlBuild.WriteString("/")
	urlBuild.WriteString(method)
	urlBuild.WriteString("/")

	req.SetRequestURI(urlBuild.String())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.Set("X-Authentication", b.SessionKey)
	req.Header.SetMethod("POST")

	if filter != nil {
		filterBody, err := ffjson.Marshal(&filter)
		if err != nil {
			return err
		}

		req.SetBody(filterBody)
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() == 400 {
		err = ffjson.Unmarshal(resp.Body(), &bettingError)
		if err != nil {
			return err
		}

		return fmt.Errorf("Error with code - %s and string - %s", bettingError.Faultcode, bettingError.Faultstring)
	}

	err = ffjson.Unmarshal(resp.Body(), reqStruct)
	if err != nil {
		return err
	}

	return nil
}
