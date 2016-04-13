package betting

import "errors"

var ErrUnknownMatchProjection = errors.New("Unknown matchProjection value")

type MatchProjection int

const (
	MP_NO_ROLLUP              MatchProjection = iota // No rollup, return raw fragments
	MP_ROLLED_UP_BY_PRICE                            // Rollup matched amounts by distinct matched prices per side.
	MP_ROLLED_UP_BY_AVG_PRICE                        // Rollup matched amounts by average matched price per side
)

var matchProjectionItems = [...]string{
	"NO_ROLLUP",
	"ROLLED_UP_BY_PRICE",
	"ROLLED_UP_BY_AVG_PRICE",
}

var matchProjectionMap = map[string]MatchProjection{
	matchProjectionItems[MP_NO_ROLLUP]:              MP_NO_ROLLUP,
	matchProjectionItems[MP_ROLLED_UP_BY_PRICE]:     MP_ROLLED_UP_BY_PRICE,
	matchProjectionItems[MP_ROLLED_UP_BY_AVG_PRICE]: MP_ROLLED_UP_BY_AVG_PRICE,
}

func (mp MatchProjection) String() string {
	return matchProjectionItems[mp]
}

func (code *MatchProjection) Check(enum string) error {
	val, ok := matchProjectionMap[enum]
	if !ok {
		return ErrUnknownMatchProjection
	}

	*code = val

	return nil
}
