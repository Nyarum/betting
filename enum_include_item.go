package betting

var EIncludeItem = eIncludeItem{
	ALL:                  "ALL",
	DEPOSITS_WITHDRAWALS: "DEPOSITS_WITHDRAWALS",
	EXCHANGE:             "EXCHANGE",
	POKER_ROOM:           "POKER_ROOM",
}

type eIncludeItem struct {
	ALL                  eIncludeItemInternal
	DEPOSITS_WITHDRAWALS eIncludeItemInternal
	EXCHANGE             eIncludeItemInternal
	POKER_ROOM           eIncludeItemInternal
}

type eIncludeItemInternal string
