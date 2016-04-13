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

var sideMap = map[string]Side{
	runnerStatusItems[BACK]: BACK,
	runnerStatusItems[LAY]:  LAY,
}

func (code Side) String() string {
	return sideItems[code]
}

func (code *Side) Check(enum string) error {
	val, ok := sideMap[enum]
	if !ok {
		return ErrUnknownSide
	}

	*code = val

	return nil
}
