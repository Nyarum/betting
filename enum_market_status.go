package betting

import "errors"

var ErrUnknowMarketStatus = errors.New("Unknown marketStatus value")

type MarketStatus int

const (
	MS_INACTIVE  MarketStatus = iota // Inactive Market
	MS_OPEN                          // Open Market
	MS_SUSPENDED                     // Suspended Market
	MS_CLOSED                        // Closed Market
)

var marketStatusItems = [...]string{
	"INACTIVE",
	"OPEN",
	"SUSPENDED",
	"CLOSED",
}

var marketStatusMap = map[string]MarketStatus{
	marketStatusItems[MS_INACTIVE]:  MS_INACTIVE,
	marketStatusItems[MS_OPEN]:      MS_OPEN,
	marketStatusItems[MS_SUSPENDED]: MS_SUSPENDED,
	marketStatusItems[MS_CLOSED]:    MS_CLOSED,
}

func (ms MarketStatus) String() string {
	return marketStatusItems[ms]
}

func (code *MarketStatus) Check(enum string) error {
	val, ok := marketStatusMap[enum]
	if !ok {
		return ErrUnknowMarketStatus
	}

	*code = val

	return nil
}
