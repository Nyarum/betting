package betting

import (
	"errors"
)

var ErrUnknownMarketBettingType = errors.New("Unknown marketBettingType value")

type MarketBettingType int
type MarketBettingTypes []MarketBettingType

const (
	MBT_ODDS                       MarketBettingType = iota // Odds Market - Any market that doesn't fit any any of the below categories.
	MBT_LINE                                                // Line Market - Now Deprecated
	MBT_RANGE                                               // Range Market - Now Deprecated
	MBT_ASIAN_HANDICAP_DOUBLE_LINE                          // Asian Handicap Market - A traditional Asian handicap market. Can be identified by marketType ASIAN_HANDICAP
	MBT_ASIAN_HANDICAP_SINGLE_LINE                          // Asian Single Line Market - A market in which there can be 0 or multiple winners. e,.g marketType TOTAL_GOALS
	// Sportsbook Odds Market. This type is deprecated and will be removed
	// in future releases, when Sportsbook markets will be represented as ODDS market but with a different product type.
	MBT_FIXED_ODDS
)

var marketBettingTypeItems = []string{
	"ODDS",
	"LINE",
	"RANGE",
	"ASIAN_HANDICAP_DOUBLE_LINE",
	"ASIAN_HANDICAP_SINGLE_LINE",
	"FIXED_ODDS",
}

var marketBettingTypeMap = map[string]MarketBettingType{
	marketBettingTypeItems[MBT_ODDS]:                       MBT_ODDS,
	marketBettingTypeItems[MBT_LINE]:                       MBT_LINE,
	marketBettingTypeItems[MBT_RANGE]:                      MBT_RANGE,
	marketBettingTypeItems[MBT_ASIAN_HANDICAP_DOUBLE_LINE]: MBT_ASIAN_HANDICAP_DOUBLE_LINE,
	marketBettingTypeItems[MBT_ASIAN_HANDICAP_SINGLE_LINE]: MBT_ASIAN_HANDICAP_SINGLE_LINE,
	marketBettingTypeItems[MBT_FIXED_ODDS]:                 MBT_FIXED_ODDS,
}

func (code MarketBettingType) String() string {
	return marketBettingTypeItems[code]
}

func (code *MarketBettingType) Check(enum string) error {
	val, ok := marketBettingTypeMap[enum]
	if !ok {
		return ErrUnknownMarketBettingType
	}

	*code = val

	return nil
}
