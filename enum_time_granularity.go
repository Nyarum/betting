package betting

import "errors"

var (
	ETimeGranularity = eTimeGranularity{
		TG_DAYS:    "DAYS",
		TG_HOURS:   "HOURS",
		TG_MINUTES: "MINUTES",
	}
	ErrUnknownTimeGranularity = errors.New("Unknown timeGranularity value")
)

type eTimeGranularity struct {
	TG_DAYS    eTimeGranularityInternal
	TG_HOURS   eTimeGranularityInternal
	TG_MINUTES eTimeGranularityInternal
}

type eTimeGranularityInternal string
