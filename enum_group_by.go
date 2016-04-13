package betting

import "errors"

var ErrUnknownGroupBy = errors.New("Unknown groupBy value")

type GroupBy int

const (
	GB_EVENT_TYPE GroupBy = iota // A roll up of settled P&L, commission paid and number of bet orders, on a specified event type
	GB_EVENT                     // A roll up of settled P&L, commission paid and number of bet orders, on a specified event
	GB_MARKET                    // A roll up of settled P&L, commission paid and number of bet orders, on a specified market
	GB_SIDE                      // An averaged roll up of settled P&L, and number of bets, on the specified side of a specified selection within a specified market, that are either settled or voided
	GB_BET                       // The P&L, commission paid, side and regulatory information etc, about each individual bet order
)

var groupByItems = [...]string{
	"EVENT_TYPE",
	"EVENT",
	"MARKET",
	"SIDE",
	"BET",
}

var groupByMap = map[string]GroupBy{
	groupByItems[GB_EVENT_TYPE]: GB_EVENT_TYPE,
	groupByItems[GB_EVENT]:      GB_EVENT,
	groupByItems[GB_MARKET]:     GB_MARKET,
	groupByItems[GB_SIDE]:       GB_SIDE,
	groupByItems[GB_BET]:        GB_BET,
}

func (code GroupBy) String() string {
	return groupByItems[code]
}

func (code *GroupBy) Check(enum string) error {
	val, ok := groupByMap[enum]
	if !ok {
		return ErrUnknownGroupBy
	}

	*code = val

	return nil
}
