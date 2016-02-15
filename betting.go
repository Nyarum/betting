package betting

import (
	"bytes"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// BetURL type of all betfair urls
type BetURL string

const (
	CertURL           = "https://identitysso-api.betfair.com/api/certlogin"
	AccountURL BetURL = "https://api.betfair.com/exchange/account/json-rpc/v1"
	BettingURL BetURL = "https://api.betfair.com/exchange/betting/json-rpc/v1"
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

// Request function for send requests to betfair via JSON RPC
func (b *Betting) Request(reqStruct interface{}, url BetURL, method string) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	req.SetRequestURI(string(url))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.Set("X-Authentication", b.SessionKey)
	req.Header.SetMethod("POST")

	bufferString := bytes.NewBuffer([]byte{})
	bufferString.WriteString(`{"params":{"filter":{"eventTypeIds":[1]}},"jsonrpc":"2.0","method":"`)
	bufferString.WriteString(method)
	bufferString.WriteString(`","id":1}`)

	req.SetBody(bufferString.Bytes())

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), reqStruct)
	if err != nil {
		return err
	}

	return nil
}
