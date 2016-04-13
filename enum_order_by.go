package betting

import "errors"

var ErrUnknownOrderBy = errors.New("Unknown orderBy value")

type OrderBy int

const (
	OB_BY_BET          OrderBy = iota // @Deprecated Use BY_PLACE_TIME instead. Order by placed time, then bet id.
	OB_BY_MARKET                      // Order by market id, then placed time, then bet id.
	OB_BY_MATCH_TIME                  // Order by time of last matched fragment (if any), then placed time, then bet id. Filters out orders which have no matched date. The dateRange filter (if specified) is applied to the matched date.
	OB_BY_PLACE_TIME                  // Order by placed time, then bet id. This is an alias of to be deprecated BY_BET. The dateRange filter (if specified) is applied to the placed date.
	OB_BY_SETTLED_TIME                // Order by time of last settled fragment (if any due to partial market settlement), then by last match time, then placed time, then bet id. Filters out orders which have not been settled. The dateRange filter (if specified) is applied to the settled date.
	OB_BY_VOID_TIME                   // Order by time of last voided fragment (if any), then by last match time, then placed time, then bet id. Filters out orders which have not been voided. The dateRange filter (if specified) is applied to the voided date.
)

var orderByItems = [...]string{
	"BY_BET",
	"BY_MARKET",
	"BY_MATCH_TIME",
	"BY_PLACE_TIME",
	"SETTLED_TIME",
	"BY_VOID_TIME",
}

var orderByMap = map[string]OrderBy{
	orderByItems[OB_BY_BET]:          OB_BY_BET,
	orderByItems[OB_BY_MARKET]:       OB_BY_MARKET,
	orderByItems[OB_BY_MATCH_TIME]:   OB_BY_MATCH_TIME,
	orderByItems[OB_BY_PLACE_TIME]:   OB_BY_PLACE_TIME,
	orderByItems[OB_BY_SETTLED_TIME]: OB_BY_SETTLED_TIME,
	orderByItems[OB_BY_VOID_TIME]:    OB_BY_VOID_TIME,
}

func (code OrderBy) String() string {
	return orderByItems[code]
}

func (code *OrderBy) Check(enum string) error {
	val, ok := orderByMap[enum]
	if !ok {
		return ErrUnknownOrderBy
	}

	*code = val

	return nil
}
