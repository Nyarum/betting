package betting

import (
	"errors"
)

var (
	ESide = eSide{
		BACK: "BACK",
		LAY:  "LAY",
	}
	ErrUnknownSide = errors.New("Unknown side value. Supported BACK or LAY only")
)

type eSide struct {
	BACK eSideInternal
	LAY  eSideInternal
}

type eSideInternal string
