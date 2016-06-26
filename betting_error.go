package betting

var bettingError = BettingError{}

type BettingError struct {
	Detail      interface{} `json:"detail"`
	Faultcode   string      `json:"faultcode"`
	Faultstring string      `json:"faultstring"`
}
