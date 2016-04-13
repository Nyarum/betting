package betting

import "errors"

var ErrUnknownExecutionReportErrorCode = errors.New("Unknown executionReportErrorCode value")

type ExecutionReportErrorCode int

const (
	EREC_ERROR_IN_MATCHER            ExecutionReportErrorCode = iota // The matcher is not healthy
	EREC_PROCESSED_WITH_ERRORS                                       // The order itself has been accepted, but at least one (possibly all) actions have generated errors
	EREC_BET_ACTION_ERROR                                            // There is an error with an action that has caused the entire order to be rejected
	EREC_INVALID_ACCOUNT_STATE                                       // Order rejected due to the account's status (suspended, inactive, dup cards)
	EREC_INVALID_WALLET_STATUS                                       // Order rejected due to the account's wallet's status
	EREC_INSUFFICIENT_FUNDS                                          // Account has exceeded its exposure limit or available to bet limit
	EREC_LOSS_LIMIT_EXCEEDED                                         // The account has exceed the self imposed loss limit
	EREC_MARKET_SUSPENDED                                            // Market is suspended
	EREC_MARKET_NOT_OPEN_FOR_BETTING                                 // Market is not open for betting, either inactive, suspended or closed
	EREC_DUPLICATE_TRANSACTION                                       // Duplicate customer reference data submitted - Please note: There is a time window associated with the de-duplication of duplicate submissions which is 60 second
	EREC_INVALID_ORDER                                               // Order cannot be accepted by the matcher due to the combination of actions. For example, bets being edited are not on the same market, or order includes both edits and placement
	EREC_INVALID_MARKET_ID                                           // Market doesn't exist
	EREC_PERMISSION_DENIED                                           // Business rules do not allow order to be placed. You are either attempting to place the order using a Delayed Application Key or from a restricted jurisdiction (i.e. USA)
	EREC_DUPLICATE_BETIDS                                            // duplicate bet ids found
	EREC_NO_ACTION_REQUIRED                                          // Order hasn't been passed to matcher as system detected there will be no state change
	EREC_SERVICE_UNAVAILABLE                                         // The requested service is unavailable
	EREC_REJECTED_BY_REGULATOR                                       // The regulator rejected the order. On the Italian Exchange this error will occur if more than 50 bets are sent in a single placeOrders request.
	EREC_NO_CHASING                                                  // A specific error code that relates to Spanish Exchange markets only which indicates that the bet placed contravenes the Spanish regulatory rules relating to loss chasing.
)

var erecItems = [...]string{
	"ERROR_IN_MATCHER",
	"PROCESSED_WITH_ERRORS",
	"BET_ACTION_ERROR",
	"INVALID_ACCOUNT_STATE",
	"INVALID_WALLET_STATUS",
	"INSUFFICIENT_FUNDS",
	"LOSS_LIMIT_EXCEEDED",
	"MARKET_SUSPENDED",
	"MARKET_NOT_OPEN_FOR_BETTING",
	"DUPLICATE_TRANSACTION",
	"INVALID_ORDER",
	"INVALID_MARKET_ID",
	"PERMISSION_DENIED",
	"DUPLICATE_BETIDS",
	"NO_ACTION_REQUIRED",
	"SERVICE_UNAVAILABLE",
	"REJECTED_BY_REGULATOR",
	"EREC_NO_CHASING",
}

var erecMap = map[string]ExecutionReportErrorCode{
	erecItems[EREC_ERROR_IN_MATCHER]:            EREC_ERROR_IN_MATCHER,
	erecItems[EREC_PROCESSED_WITH_ERRORS]:       EREC_PROCESSED_WITH_ERRORS,
	erecItems[EREC_BET_ACTION_ERROR]:            EREC_BET_ACTION_ERROR,
	erecItems[EREC_INVALID_ACCOUNT_STATE]:       EREC_INVALID_ACCOUNT_STATE,
	erecItems[EREC_INVALID_WALLET_STATUS]:       EREC_INVALID_WALLET_STATUS,
	erecItems[EREC_INSUFFICIENT_FUNDS]:          EREC_INSUFFICIENT_FUNDS,
	erecItems[EREC_LOSS_LIMIT_EXCEEDED]:         EREC_LOSS_LIMIT_EXCEEDED,
	erecItems[EREC_MARKET_SUSPENDED]:            EREC_MARKET_SUSPENDED,
	erecItems[EREC_MARKET_NOT_OPEN_FOR_BETTING]: EREC_MARKET_NOT_OPEN_FOR_BETTING,
	erecItems[EREC_DUPLICATE_TRANSACTION]:       EREC_DUPLICATE_TRANSACTION,
	erecItems[EREC_INVALID_ORDER]:               EREC_INVALID_ORDER,
	erecItems[EREC_INVALID_MARKET_ID]:           EREC_INVALID_MARKET_ID,
	erecItems[EREC_PERMISSION_DENIED]:           EREC_PERMISSION_DENIED,
	erecItems[EREC_DUPLICATE_BETIDS]:            EREC_DUPLICATE_BETIDS,
	erecItems[EREC_NO_ACTION_REQUIRED]:          EREC_NO_ACTION_REQUIRED,
	erecItems[EREC_SERVICE_UNAVAILABLE]:         EREC_SERVICE_UNAVAILABLE,
	erecItems[EREC_REJECTED_BY_REGULATOR]:       EREC_REJECTED_BY_REGULATOR,
}

func (code ExecutionReportErrorCode) String() string {
	return erecItems[code]
}

func (code *ExecutionReportErrorCode) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := erecMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownExecutionReportErrorCode
	}
	return err

}

var ErrUnknownExecutionReportStatus = errors.New("Unknown executionReportStatus value")

type ExecutionReportStatus int

const (
	ERS_SUCCESS ExecutionReportStatus = iota // Order processed successfully
	ERS_FAILURE                              // Order failed.

	// The order itself has been accepted, but at least one (possibly all) actions have generated errors.
	// This error only occurs for replaceOrders, cancelOrders and updateOrders operations.
	// The placeOrders operation will not return PROCESSED_WITH_ERRORS status as it is an atomic operation.
	ERS_PROCESSED_WITH_ERRORS
	ERS_TIMEOUT // Order timed out.
)

var ersItems = [...]string{
	"SUCCESS",
	"FAILURE",
	"PROCESSED_WITH_ERRORS",
	"TIMEOUT",
}

var ersMap = map[string]ExecutionReportStatus{
	ersItems[ERS_SUCCESS]:               ERS_SUCCESS,
	ersItems[ERS_FAILURE]:               ERS_FAILURE,
	ersItems[ERS_PROCESSED_WITH_ERRORS]: ERS_PROCESSED_WITH_ERRORS,
	ersItems[ERS_TIMEOUT]:               ERS_TIMEOUT,
}

func (code ExecutionReportStatus) String() string {
	return ersItems[code]
}

func (code *ExecutionReportStatus) Check(enum string) error {
	val, ok := ersMap[enum]
	if !ok {
		return ErrUnknownExecutionReportStatus
	}

	*code = val

	return nil
}
