package betting

import "errors"

var ErrUnknownRollupModel = errors.New("Unknown rollupModel value")

type RollupModel int

const (
	RM_STAKE             RollupModel = iota // The volumes will be rolled up to the minimum value which is >= rollupLimit.
	RM_PAYOUT                               // The volumes will be rolled up to the minimum value where the payout( price * volume ) is >= rollupLimit.
	RM_MANAGED_LIABILITY                    // The volumes will be rolled up to the minimum value which is >= rollupLimit, until a lay price threshold. There after, the volumes will be rolled up to the minimum value such that the liability >= a minimum liability. Not supported as yet.
	RM_NONE                                 // No rollup will be applied. However the volumes will be filtered by currency specific minimum stake unless overridden specifically for the channel.
)

var rollupModelItems = [...]string{
	"STAKE",
	"PAYOUT",
	"MANAGED_LIABILITY",
	"NONE",
}

var rollupModelMap = map[string]RollupModel{
	rollupModelItems[RM_STAKE]:             RM_STAKE,
	rollupModelItems[RM_PAYOUT]:            RM_PAYOUT,
	rollupModelItems[RM_MANAGED_LIABILITY]: RM_MANAGED_LIABILITY,
	rollupModelItems[RM_NONE]:              RM_NONE,
}

func (code RollupModel) String() string {
	return rollupModelItems[code]
}

func (code *RollupModel) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := rollupModelMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownRollupModel
	}
	return err

}
