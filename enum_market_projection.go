package betting

import (
	"errors"
)

var ErrUnknownMarketProjection = errors.New("Unknown marketProjection value")

type MarketProjection int
type MarketProjections []MarketProjection

const (
	MP_COMPETITION        MarketProjection = iota // If not selected then the competition will not be returned with marketCatalogue
	MP_EVENT                                      // If not selected then the event will not be returned with marketCatalogue
	MP_EVENT_TYPE                                 // If not selected then the eventType will not be returned with marketCatalogue
	MP_MARKET_START_TIME                          // If not selected then the start time will not be returned with marketCatalogue
	MP_MARKET_DESCRIPTION                         // If not selected then the description will not be returned with marketCatalogue
	MP_RUNNER_DESCRIPTION                         // If not selected then the runners will not be returned with marketCatalogue
	MP_RUNNER_METADATA                            // If not selected then the runner metadata will not be returned with marketCatalogue. If selected then RUNNER_DESCRIPTION will also be returned regardless of whether it is included as a market projection.
)

var marketProjectionItems = [...]string{
	"COMPETITION",
	"EVENT",
	"EVENT_TYPE",
	"MARKET_START_TIME",
	"MARKET_DESCRIPTION",
	"RUNNER_DESCRIPTION",
	"RUNNER_METADATA",
}

var marketProjectionMap = map[string]MarketProjection{
	marketProjectionItems[MP_COMPETITION]:        MP_COMPETITION,
	marketProjectionItems[MP_EVENT]:              MP_EVENT,
	marketProjectionItems[MP_EVENT_TYPE]:         MP_EVENT_TYPE,
	marketProjectionItems[MP_MARKET_START_TIME]:  MP_MARKET_START_TIME,
	marketProjectionItems[MP_MARKET_DESCRIPTION]: MP_MARKET_DESCRIPTION,
	marketProjectionItems[MP_RUNNER_DESCRIPTION]: MP_RUNNER_DESCRIPTION,
	marketProjectionItems[MP_RUNNER_METADATA]:    MP_RUNNER_METADATA,
}

func (code MarketProjection) String() string {
	return marketProjectionItems[code]
}

func (code *MarketProjection) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := marketProjectionMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownMarketProjection
	}
	return err

}
