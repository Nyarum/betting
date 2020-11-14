package betting

// BetfairRestURL type of all betfair rest urls
type BetfairRestURL string

type BetfairRestURLs struct {
	AccountURL   BetfairRestURL
	BettingURL   BetfairRestURL
	CertURL      BetfairRestURL
	KeepAliveURL BetfairRestURL
	ScoresURL    BetfairRestURL
}

type Betfair struct {
	*Client
	*Betting
}

func NewBetfair(apikey string) *Betfair {
	var defaultAccountURL BetfairRestURL = "https://api.betfair.com/exchange/account/rest/v1.0"
	var defaultBettingURL BetfairRestURL = "https://api.betfair.com/exchange/betting/rest/v1.0"
	var defaultCertURL BetfairRestURL = "https://identitysso-cert.betfair.com/api/certlogin"
	var defaultKeepAliveURL BetfairRestURL = "https://identitysso.betfair.com/api/keepAlive"
	var defaultScoresURL BetfairRestURL = "https://api.betfair.com/exchange/scores/json-rpc/v1"

	return NewBetfairWithSpecifiedURLs(apikey, BetfairRestURLs{
		AccountURL:   defaultAccountURL,
		BettingURL:   defaultBettingURL,
		CertURL:      defaultCertURL,
		KeepAliveURL: defaultKeepAliveURL,
		ScoresURL:    defaultScoresURL,
	})
}

func NewBetfairWithSpecifiedURLs(apikey string, urls BetfairRestURLs) *Betfair {
	client := Client{
		ApiKey:       apikey,
		CertURL:      urls.CertURL,
		KeepAliveURL: urls.KeepAliveURL,
	}

	return &Betfair{
		Client: &client,
		Betting: &Betting{
			Client:     &client,
			BettingURL: urls.BettingURL,
		},
	}
}

// TODO: Deprecate
var NewBet = NewBetfair
