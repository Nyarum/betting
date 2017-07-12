package betting

type EEventStatus string

const (
	ES_FINISHED    EEventStatus = "FINISHED"
	ES_IN_PROGRESS EEventStatus = "IN_PROGRESS"
	ES_PENDING     EEventStatus = "PENDING"
)

type EResponseCode string

const (
	RC_OK                                EResponseCode = "OK"
	RC_NO_NEW_UPDATES                    EResponseCode = "NO_NEW_UPDATES"
	RC_NO_LIVE_DATA_AVAILABLE            EResponseCode = "NO_LIVE_DATA_AVAILABLE"
	RC_SERVICE_UNAVAILABLE               EResponseCode = "SERVICE_UNAVAILABLE"
	RC_UNEXPECTED_ERROR                  EResponseCode = "UNEXPECTED_ERROR"
	RC_LIVE_DATA_TEMPORARILY_UNAVAILABLE EResponseCode = "LIVE_DATA_TEMPORARILY_UNAVAILABLE"
)
