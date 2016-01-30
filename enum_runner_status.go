package betting

import "errors"

var ErrUnknownRunnerStatus = errors.New("Unknown runnerStatus value")

type RunnerStatus int

const (
	RS_ACTIVE         RunnerStatus = iota // ACTIVE
	RS_WINNER                             // WINNER
	RS_LOSER                              // LOSER
	RS_REMOVED_VACANT                     // REMOVED_VACANT applies to Greyhounds. Greyhound markets always return a fixed number of runners (traps). If a dog has been removed, the trap is shown as vacant.
	RS_REMOVED                            // REMOVED
	// The selection is hidden from the market.  This occurs in Horse Racing markets were runners is hidden
	// when it is doesn’t hold an official entry following an entry stage.
	// This could be because the horse was never entered or because they have been scratched from a race at a declaration stage.
	// All matched customer bet prices are set to 1.0 even if there are later supplementary stages.
	// Should it appear likely that a specific runner may actually be supplemented into the race this runner will be reinstated
	// with all matched customer bets set back to the original prices.
	RS_HIDDEN
)

var runnerStatusItems = [...]string{
	"ACTIVE",
	"WINNER",
	"LOSER",
	"REMOVED_VACANT",
	"REMOVED",
	"HIDDEN",
}

var runnerStatusMap = map[string]RunnerStatus{
	runnerStatusItems[RS_ACTIVE]:         RS_ACTIVE,
	runnerStatusItems[RS_WINNER]:         RS_WINNER,
	runnerStatusItems[RS_LOSER]:          RS_LOSER,
	runnerStatusItems[RS_REMOVED_VACANT]: RS_REMOVED_VACANT,
	runnerStatusItems[RS_REMOVED]:        RS_REMOVED,
	runnerStatusItems[RS_HIDDEN]:         RS_HIDDEN,
}

func (ms RunnerStatus) String() string {
	return runnerStatusItems[ms]
}

func (code *RunnerStatus) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := runnerStatusMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownRunnerStatus
	}
	return err

}
