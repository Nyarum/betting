package betting

import "errors"

var (
	EBetStatus = eBetStatus{
		BS_SETTLED:   "SETTLED",
		BS_VOIDED:    "VOIDED",
		BS_LAPSED:    "LAPSED",
		BS_CANCELLED: "CANCELLED",
	}
	ErrUnknownBetStatus = errors.New("Unknown betStatus value")
)

type eBetStatus struct {
	BS_SETTLED   eBetStatusInternal
	BS_VOIDED    eBetStatusInternal
	BS_LAPSED    eBetStatusInternal
	BS_CANCELLED eBetStatusInternal
}

type eBetStatusInternal string
