package betting

import "errors"

var (
	EOrderBy = eOrderBy{
		OB_BY_BET:          "BY_BET",
		OB_BY_MARKET:       "BY_MARKET",
		OB_BY_MATCH_TIME:   "BY_MATCH_TIME",
		OB_BY_PLACE_TIME:   "BY_PLACE_TIME",
		OB_BY_SETTLED_TIME: "BY_SETTLED_TIME",
		OB_BY_VOID_TIME:    "BY_VOID_TIME",
	}
	ErrUnknownOrderBy = errors.New("Unknown orderBy value")
)

type eOrderBy struct {
	OB_BY_BET          eOrderByInternal
	OB_BY_MARKET       eOrderByInternal
	OB_BY_MATCH_TIME   eOrderByInternal
	OB_BY_PLACE_TIME   eOrderByInternal
	OB_BY_SETTLED_TIME eOrderByInternal
	OB_BY_VOID_TIME    eOrderByInternal
}

type eOrderByInternal string
