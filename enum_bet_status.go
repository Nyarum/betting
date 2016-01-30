package betting

import "errors"

var ErrUnknownBetStatus = errors.New("Unknown betStatus value")

type BetStatus int

const (
	BS_SETTLED   BetStatus = iota // A matched bet that was settled normally
	BS_VOIDED                     // A matched bet that was subsequently voided by Betfair, before, during or after settlement
	BS_LAPSED                     // Unmatched bet that was cancelled by Betfair (for example at turn in play).
	BS_CANCELLED                  // Unmatched bet that was cancelled by an explicit customer action.
)

var betStatusItems = [...]string{
	"SETTLED",
	"VOIDED",
	"LAPSED",
	"CANCELLED",
}

var betStatusMap = map[string]BetStatus{
	betStatusItems[BS_SETTLED]:   BS_SETTLED,
	betStatusItems[BS_VOIDED]:    BS_VOIDED,
	betStatusItems[BS_LAPSED]:    BS_LAPSED,
	betStatusItems[BS_CANCELLED]: BS_CANCELLED,
}

func (code BetStatus) String() string {
	return betStatusItems[code]
}

func (code *BetStatus) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := betStatusMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownBetStatus
	}
	return err

}
