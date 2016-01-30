package betting

import "errors"

var ErrUnknownExceptionCode = errors.New("Unknown APINGException errorCode value")

// ExceptionErrorCode holds code of thrown exception when an operation fails
type ExceptionErrorCode int

const (
	EEC_TOO_MUCH_DATA               ExceptionErrorCode = iota // The operation requested too much data, exceeding the Market Data Request Limits.
	EEC_INVALID_INPUT_DATA                                    // The data input is invalid. A specific description is returned via errorDetails as shown below.
	EEC_INVALID_SESSION_INFORMATION                           // The session token hasn't been provided, is invalid or has expired.
	EEC_NO_APP_KEY                                            // An application key header ('X-Application') has not been provided in the request
	EEC_NO_SESSION                                            // A session token header ('X-Authentication') has not been provided in the request
	EEC_UNEXPECTED_ERROR                                      // An unexpected internal error occurred that prevented successful request processing.
	EEC_INVALID_APP_KEY                                       // The application key passed is invalid or is not present
	EEC_TOO_MANY_REQUESTS                                     // There are too many pending requests e.g. a listMarketBook with Order/Match projections is limited to 3 concurrent requests. The error also applies to listCurrentOrders, listMarketProfitAndLoss and listClearedOrders if you have 3 or more requests currently in execution
	EEC_SERVICE_BUSY                                          // The service is currently too busy to service this request.
	EEC_TIMEOUT_ERROR                                         // The Internal call to downstream service timed out. Please note: If a TIMEOUT_ERROR error occurs on a placeOrders/replaceOrders request, you should check listCurrentOrders to verify the status of your bets before placing further orders. Please allow up to 2 minutes for timed out order to appear.
	EEC_REQUEST_SIZE_EXCEEDS_LIMIT                            // The request exceeds the request size limit. Requests are limited to a total of 250 betId’s/marketId’s (or a combination of both).
	EEC_ACCESS_DENIED                                         // The calling client is not permitted to perform the specific action e.g. the using a Delayed App Key when placing bets or attempting to place a bet from a restricted jurisdiction.
)

var eecItems = [...]string{
	"TOO_MUCH_DATA",
	"INVALID_INPUT_DATA",
	"INVALID_SESSION_INFORMATION",
	"NO_APP_KEY",
	"NO_SESSION",
	"UNEXPECTED_ERROR",
	"INVALID_APP_KEY",
	"TOO_MANY_REQUESTS",
	"SERVICE_BUSY",
	"TIMEOUT_ERROR",
	"REQUEST_SIZE_EXCEEDS_LIMIT",
	"ACCESS_DENIED",
}

var eecMap = map[string]ExceptionErrorCode{
	eecItems[EEC_TOO_MUCH_DATA]:               EEC_TOO_MUCH_DATA,
	eecItems[EEC_INVALID_INPUT_DATA]:          EEC_INVALID_INPUT_DATA,
	eecItems[EEC_INVALID_SESSION_INFORMATION]: EEC_INVALID_SESSION_INFORMATION,
	eecItems[EEC_NO_APP_KEY]:                  EEC_NO_APP_KEY,
	eecItems[EEC_NO_SESSION]:                  EEC_NO_SESSION,
	eecItems[EEC_UNEXPECTED_ERROR]:            EEC_UNEXPECTED_ERROR,
	eecItems[EEC_INVALID_APP_KEY]:             EEC_INVALID_APP_KEY,
	eecItems[EEC_TOO_MANY_REQUESTS]:           EEC_TOO_MANY_REQUESTS,
	eecItems[EEC_SERVICE_BUSY]:                EEC_SERVICE_BUSY,
	eecItems[EEC_TIMEOUT_ERROR]:               EEC_TIMEOUT_ERROR,
	eecItems[EEC_REQUEST_SIZE_EXCEEDS_LIMIT]:  EEC_REQUEST_SIZE_EXCEEDS_LIMIT,
	eecItems[EEC_ACCESS_DENIED]:               EEC_ACCESS_DENIED,
}

func (code ExceptionErrorCode) String() string {
	return eecItems[code]
}

func (code *ExceptionErrorCode) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := eecMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownExceptionCode
	}
	return err
}
