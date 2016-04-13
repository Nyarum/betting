package betting

import "errors"

var ErrUnknownMarketSort = errors.New("Unknown marketSort value")

type MarketSort int

const (
	MS_MINIMUM_TRADED    MarketSort = iota // Minimum traded volume
	MS_MAXIMUM_TRADED                      // Maximum traded volume
	MS_MINIMUM_AVAILABLE                   // Minimum available to match
	MS_MAXIMUM_AVAILABLE                   // Maximum available to match
	MS_FIRST_TO_START                      // The closest markets based on their expected start time
	MS_LAST_TO_START                       // The most distant markets based on their expected start time
)

var marketSortItems = [...]string{
	"MINIMUM_TRADED",
	"MAXIMUM_TRADED",
	"MINIMUM_AVAILABLE",
	"MAXIMUM_AVAILABLE",
	"FIRST_TO_START",
	"LAST_TO_START",
}

var marketSortMap = map[string]MarketSort{
	marketSortItems[MS_MINIMUM_TRADED]:    MS_MINIMUM_TRADED,
	marketSortItems[MS_MINIMUM_TRADED]:    MS_MINIMUM_TRADED,
	marketSortItems[MS_MINIMUM_AVAILABLE]: MS_MINIMUM_AVAILABLE,
	marketSortItems[MS_MAXIMUM_AVAILABLE]: MS_MAXIMUM_AVAILABLE,
	marketSortItems[MS_FIRST_TO_START]:    MS_FIRST_TO_START,
	marketSortItems[MS_LAST_TO_START]:     MS_LAST_TO_START,
}

func (code MarketSort) String() string {
	return marketSortItems[code]
}

func (code *MarketSort) Check(enum string) error {
	val, ok := marketSortMap[enum]
	if !ok {
		return ErrUnknownMarketSort
	}

	*code = val

	return nil
}
