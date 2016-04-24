package betting

type IncludeItem int

const (
	ALL IncludeItem = iota
	DEPOSITS_WITHDRAWALS
	EXCHANGE
	POKER_ROOM
)
