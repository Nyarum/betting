package betting

import "errors"

var ErrUnknownOrderType = errors.New("Unknown orderType value")

type OrderType int

const (
	OT_LIMIT           OrderType = iota // A normal exchange limit order for immediate execution
	OT_LIMIT_ON_CLOSE                   // Limit order for the auction (SP)
	OT_MARKET_ON_CLOSE                  // Market order for the auction (SP)
)

var orderTypeItems = [...]string{
	"BY_BET",
	"BY_MARKET",
	"BY_MATCH_TIME",
	"BY_PLACE_TIME",
	"BY_SETTLED_TIME",
	"BY_VOID_TIME",
}

var orderTypeMap = map[string]OrderType{
	orderTypeItems[OT_LIMIT]:           OT_LIMIT,
	orderTypeItems[OT_LIMIT_ON_CLOSE]:  OT_LIMIT_ON_CLOSE,
	orderTypeItems[OT_MARKET_ON_CLOSE]: OT_MARKET_ON_CLOSE,
}

func (code OrderType) String() string {
	return orderTypeItems[code]
}

func (code *OrderType) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := orderTypeMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownOrderType
	}
	return err

}
