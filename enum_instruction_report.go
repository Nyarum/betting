package betting

import "errors"

var ErrUnknownInstructionReportStatus = errors.New("Unknown instructionReportStatus value")

type InstructionReportStatus int

const (
	IRS_SUCCESS InstructionReportStatus = iota
	IRS_FAILURE
	IRS_TIMEOUT
)

var instructionReportStatusItems = [...]string{
	"SUCCESS",
	"FAILURE",
	"TIMEOUT",
}

var instructionReportStatusMap = map[string]InstructionReportStatus{
	instructionReportStatusItems[IRS_SUCCESS]: IRS_SUCCESS,
	instructionReportStatusItems[IRS_FAILURE]: IRS_FAILURE,
	instructionReportStatusItems[IRS_TIMEOUT]: IRS_TIMEOUT,
}

func (code InstructionReportStatus) String() string {
	return instructionReportStatusItems[code]
}

func (code *InstructionReportStatus) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := instructionReportStatusMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownInstructionReportStatus
	}
	return err
}

var ErrUnknownInstructionReportErrorCode = errors.New("Unknown instructionReportErrorCode value")

type InstructionReportErrorCode int

const (
	IREC_INVALID_BET_SIZE                InstructionReportErrorCode = iota // bet size is invalid for your currency or your regulator
	IREC_INVALID_RUNNER                                                    // Runner does not exist, includes vacant traps in greyhound racing
	IREC_BET_TAKEN_OR_LAPSED                                               // Bet cannot be cancelled or modified as it has already been taken or has lapsed Includes attempts to cancel/modify market on close BSP bets and cancelling limit on close BSP bets
	IREC_BET_IN_PROGRESS                                                   // No result was received from the matcher in a timeout configured for the system
	IREC_RUNNER_REMOVED                                                    // Runner has been removed from the event
	IREC_MARKET_NOT_OPEN_FOR_BETTING                                       // Attempt to edit a bet on a market that has closed.
	IREC_LOSS_LIMIT_EXCEEDED                                               // The action has caused the account to exceed the self imposed loss limit
	IREC_MARKET_NOT_OPEN_FOR_BSP_BETTING                                   // Market now closed to bsp betting. Turned in-play or has been reconciled
	IREC_INVALID_PRICE_EDIT                                                // Attempt to edit down the price of a bsp limit on close lay bet, or edit up the price of a limit on close back bet
	IREC_INVALID_ODDS                                                      // Odds not on price ladder - either edit or placement
	IREC_INSUFFICIENT_FUNDS                                                // Insufficient funds available to cover the bet action. Either the exposure limit or available to bet limit would be exceeded
	IREC_INVALID_PERSISTENCE_TYPE                                          // Invalid persistence type for this market, e.g. KEEP for a non bsp market
	IREC_ERROR_IN_MATCHER                                                  // A problem with the matcher prevented this action completing successfully
	IREC_INVALID_BACK_LAY_COMBINATION                                      //  The order contains a back and a lay for the same runner at overlapping prices. This would guarantee a self match. This also applies to BSP limit on close bets
	IREC_ERROR_IN_ORDER                                                    // The action failed because the parent order failed
	IREC_INVALID_BID_TYPE                                                  // Bid type is mandatory
	IREC_INVALID_BET_ID                                                    // Bet for id supplied has not been found
	IREC_CANCELLED_NOT_PLACED                                              // Bet cancelled but replacement bet was not placed
	IREC_RELATED_ACTION_FAILED                                             // Action failed due to the failure of a action on which this action is dependent
	IREC_NO_ACTION_REQUIRED                                                // the action does not result in any state change. eg changing a persistence to it's current value
)

var irecItems = [...]string{
	"INVALID_BET_SIZE",
	"INVALID_RUNNER",
	"BET_TAKEN_OR_LAPSED",
	"BET_IN_PROGRESS",
	"RUNNER_REMOVED",
	"MARKET_NOT_OPEN_FOR_BETTING",
	"LOSS_LIMIT_EXCEEDED",
	"MARKET_NOT_OPEN_FOR_BSP_BETTING",
	"INVALID_PRICE_EDIT",
	"INVALID_ODDS",
	"INSUFFICIENT_FUNDS",
	"INVALID_PERSISTENCE_TYPE",
	"ERROR_IN_MATCHER",
	"INVALID_BACK_LAY_COMBINATION",
	"ERROR_IN_ORDER",
	"INVALID_BID_TYPE",
	"INVALID_BET_ID",
	"CANCELLED_NOT_PLACED",
	"RELATED_ACTION_FAILED",
	"NO_ACTION_REQUIRED",
}

var irecMap = map[string]InstructionReportErrorCode{
	irecItems[IREC_INVALID_BET_SIZE]:                IREC_INVALID_BET_SIZE,
	irecItems[IREC_INVALID_RUNNER]:                  IREC_INVALID_RUNNER,
	irecItems[IREC_BET_TAKEN_OR_LAPSED]:             IREC_BET_TAKEN_OR_LAPSED,
	irecItems[IREC_BET_IN_PROGRESS]:                 IREC_BET_IN_PROGRESS,
	irecItems[IREC_RUNNER_REMOVED]:                  IREC_RUNNER_REMOVED,
	irecItems[IREC_MARKET_NOT_OPEN_FOR_BETTING]:     IREC_MARKET_NOT_OPEN_FOR_BETTING,
	irecItems[IREC_LOSS_LIMIT_EXCEEDED]:             IREC_LOSS_LIMIT_EXCEEDED,
	irecItems[IREC_MARKET_NOT_OPEN_FOR_BSP_BETTING]: IREC_MARKET_NOT_OPEN_FOR_BSP_BETTING,
	irecItems[IREC_INVALID_PRICE_EDIT]:              IREC_INVALID_PRICE_EDIT,
	irecItems[IREC_INVALID_ODDS]:                    IREC_INVALID_ODDS,
	irecItems[IREC_INSUFFICIENT_FUNDS]:              IREC_INSUFFICIENT_FUNDS,
	irecItems[IREC_INVALID_PERSISTENCE_TYPE]:        IREC_INVALID_PERSISTENCE_TYPE,
	irecItems[IREC_ERROR_IN_MATCHER]:                IREC_ERROR_IN_MATCHER,
	irecItems[IREC_INVALID_BACK_LAY_COMBINATION]:    IREC_INVALID_BACK_LAY_COMBINATION,
	irecItems[IREC_ERROR_IN_ORDER]:                  IREC_ERROR_IN_ORDER,
	irecItems[IREC_INVALID_BID_TYPE]:                IREC_INVALID_BID_TYPE,
	irecItems[IREC_INVALID_BET_ID]:                  IREC_INVALID_BET_ID,
	irecItems[IREC_CANCELLED_NOT_PLACED]:            IREC_CANCELLED_NOT_PLACED,
	irecItems[IREC_RELATED_ACTION_FAILED]:           IREC_RELATED_ACTION_FAILED,
	irecItems[IREC_NO_ACTION_REQUIRED]:              IREC_NO_ACTION_REQUIRED,
}

func (code InstructionReportErrorCode) String() string {
	return irecItems[code]
}

func (code *InstructionReportErrorCode) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := irecMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownInstructionReportErrorCode
	}
	return err
}
