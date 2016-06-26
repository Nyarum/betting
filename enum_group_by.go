package betting

import "errors"

var (
	EGroupBy = eGroupBy{
		GB_EVENT_TYPE: "EVENT_TYPE",
		GB_EVENT:      "EVENT",
		GB_MARKET:     "MARKET",
		GB_SIDE:       "SIDE",
		GB_BET:        "BET",
	}
	ErrUnknownGroupBy = errors.New("Unknown groupBy value")
)

type eGroupBy struct {
	GB_EVENT_TYPE eGroupByInternal
	GB_EVENT      eGroupByInternal
	GB_MARKET     eGroupByInternal
	GB_SIDE       eGroupByInternal
	GB_BET        eGroupByInternal
}

type eGroupByInternal string
