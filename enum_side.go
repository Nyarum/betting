package betting

import (
	"errors"
)

var ErrUnknownSide = errors.New("Unknown side value. Supported BACK or LAY only")

// Side represents bet selection
type Side int

// To back a team, horse or outcome is to bet on the selection to win.
// To lay a team, horse, or outcome is to bet on the selection to lose.
const (
	BACK Side = iota
	LAY
)

var sideItems = [...]string{
	"BACK",
	"LAY",
}

func (code Side) String() string {
	return sideItems[code]
}

func (side *Side) UnmarshalJSON(buf []byte) error {
	var err error
	switch {
	case string(buf) == sideItems[BACK]:
		*side = BACK
	case string(buf) == sideItems[LAY]:
		*side = LAY
	default:
		err = ErrUnknownSide
	}
	return err
}
