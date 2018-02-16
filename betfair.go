package betting

// BetfairRestURL type of all betfair rest urls
type BetfairRestURL string

const (
	CertURL                     = "https://identitysso-api.betfair.com/api/certlogin"
	KeepAliveURL                = "https://identitysso.betfair.com/api/keepAlive"
	AccountURL   BetfairRestURL = "https://api.betfair.com/exchange/account/rest/v1.0"
	BettingURL   BetfairRestURL = "https://api.betfair.com/exchange/betting/rest/v1.0"
	ScoresURL                   = "https://api.betfair.com/exchange/scores/json-rpc/v1"
)

type Betfair struct {
	*Client
	*Betting
}

func NewBetfair(apikey string) *Betfair {

	client := Client{ApiKey: apikey}

	return &Betfair{
		Client:  &client,
		Betting: &Betting{&client},
	}
}

// TODO: Deprecate
var NewBet = NewBetfair
