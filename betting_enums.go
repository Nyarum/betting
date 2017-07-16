package betting

// http://docs.developer.betfair.com/docs/display/1smk3cen4v3lu3yomq5qye0ni/Betting+Enums
// Enum version from Apr 13, 2017

type EWallet string

const (
	W_UK  EWallet = "UK"
	W_AUS EWallet = "AUSTRALIAN"
)

type EBetStatus string

const (
	BS_SETTLED          EBetStatus = "SETTLED"
	BS_VOIDED           EBetStatus = "VOIDED"
	BS_LAPSED           EBetStatus = "LAPSED"
	BS_CANCELLED        EBetStatus = "CANCELLED"
	ErrUnknownBetStatus EBetStatus = "Unknown betStatus value"
)

type EExceptionErrorCode string

const (
	EEC_TOO_MUCH_DATA               EExceptionErrorCode = "TOO_MUCH_DATA"
	EEC_INVALID_INPUT_DATA          EExceptionErrorCode = "INVALID_INPUT_DATA"
	EEC_INVALID_SESSION_INFORMATION EExceptionErrorCode = "INVALID_SESSION_INFORMATION"
	EEC_NO_APP_KEY                  EExceptionErrorCode = "NO_APP_KEY"
	EEC_NO_SESSION                  EExceptionErrorCode = "NO_SESSION"
	EEC_UNEXPECTED_ERROR            EExceptionErrorCode = "UNEXPECTED_ERROR"
	EEC_INVALID_APP_KEY             EExceptionErrorCode = "INVALID_APP_KEY"
	EEC_TOO_MANY_REQUESTS           EExceptionErrorCode = "TOO_MANY_REQUESTS"
	EEC_SERVICE_BUSY                EExceptionErrorCode = "SERVICE_BUSY"
	EEC_TIMEOUT_ERROR               EExceptionErrorCode = "TIMEOUT_ERROR"
	EEC_REQUEST_SIZE_EXCEEDS_LIMIT  EExceptionErrorCode = "REQUEST_SIZE_EXCEEDS_LIMIT"
	EEC_ACCESS_DENIED               EExceptionErrorCode = "ACCESS_DENIED"
	ErrUnknownExceptionCode         EExceptionErrorCode = "Unknown APINGException errorCode value"
)

type EExecutionReportStatus string

const (
	ERS_SUCCESS               EExecutionReportStatus = "SUCCESS"
	ERS_FAILURE               EExecutionReportStatus = "FAILURE"
	ERS_PROCESSED_WITH_ERRORS EExecutionReportStatus = "PROCESSED_WITH_ERRORS"
	ERS_TIMEOUT               EExecutionReportStatus = "TIMEOUT"
)

type EExecutionReportErrorCode string

const (
	EREC_ERROR_IN_MATCHER              EExecutionReportErrorCode = "ERROR_IN_MATCHER"
	EREC_PROCESSED_WITH_ERRORS         EExecutionReportErrorCode = "PROCESSED_WITH_ERRORS"
	EREC_BET_ACTION_ERROR              EExecutionReportErrorCode = "BET_ACTION_ERROR"
	EREC_INVALID_ACCOUNT_STATE         EExecutionReportErrorCode = "INVALID_ACCOUNT_STATE"
	EREC_INVALID_WALLET_STATUS         EExecutionReportErrorCode = "INVALID_WALLET_STATUS"
	EREC_INSUFFICIENT_FUNDS            EExecutionReportErrorCode = "INSUFFICIENT_FUNDS"
	EREC_LOSS_LIMIT_EXCEEDED           EExecutionReportErrorCode = "LOSS_LIMIT_EXCEEDED"
	EREC_MARKET_SUSPENDED              EExecutionReportErrorCode = "MARKET_SUSPENDED"
	EREC_MARKET_NOT_OPEN_FOR_BETTING   EExecutionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"
	EREC_DUPLICATE_TRANSACTION         EExecutionReportErrorCode = "DUPLICATE_TRANSACTION"
	EREC_INVALID_ORDER                 EExecutionReportErrorCode = "INVALID_ORDER"
	EREC_INVALID_MARKET_ID             EExecutionReportErrorCode = "INVALID_MARKET_ID"
	EREC_PERMISSION_DENIED             EExecutionReportErrorCode = "PERMISSION_DENIED"
	EREC_DUPLICATE_BETIDS              EExecutionReportErrorCode = "DUPLICATE_BETIDS"
	EREC_NO_ACTION_REQUIRED            EExecutionReportErrorCode = "NO_ACTION_REQUIRED"
	EREC_SERVICE_UNAVAILABLE           EExecutionReportErrorCode = "SERVICE_UNAVAILABLE"
	EREC_REJECTED_BY_REGULATOR         EExecutionReportErrorCode = "REJECTED_BY_REGULATOR"
	EREC_NO_CHASING                    EExceptionErrorCode       = "NO_CHASING"
	EREC_REGULATOR_IS_NOT_AVAILABLE    EExceptionErrorCode       = "REGULATOR_IS_NOT_AVAILABLE"
	EREC_TOO_MANY_INSTRUCTIONS         EExceptionErrorCode       = "TOO_MANY_INSTRUCTIONS"
	EREC_INVALID_MARKET_VERSION        EExceptionErrorCode       = "INVALID_MARKET_VERSION"
	ErrUnknownExecutionReportErrorCode EExecutionReportErrorCode = "Unknown executionReportErrorCode value"
)

type EGroupBy string

const (
	GB_EVENT_TYPE     EGroupBy = "EVENT_TYPE"
	GB_EVENT          EGroupBy = "EVENT"
	GB_MARKET         EGroupBy = "MARKET"
	GB_SIDE           EGroupBy = "SIDE"
	GB_BET            EGroupBy = "BET"
	GB_RUNNER         EGroupBy = "RUNNER"
	GB_STRATEGY       EGroupBy = "STRATEGY"
	ErrUnknownGroupBy EGroupBy = "Unknown groupBy value"
)

type EIncludeItem string

const (
	IT_ALL                  EIncludeItem = "ALL"
	IT_DEPOSITS_WITHDRAWALS EIncludeItem = "DEPOSITS_WITHDRAWALS"
	IT_EXCHANGE             EIncludeItem = "EXCHANGE"
	IT_POKER_ROOM           EIncludeItem = "POKER_ROOM"
)

type EInstructionReportStatus string

const (
	IRS_SUCCESS EInstructionReportStatus = "SUCCESS"
	IRS_FAILURE EInstructionReportStatus = "FAILURE"
	IRS_TIMEOUT EInstructionReportStatus = "TIMEOUT"
)

type EInstructionReportErrorCode string

const (
	IREC_INVALID_BET_SIZE                EInstructionReportErrorCode = "INVALID_BET_SIZE"
	IREC_INVALID_RUNNER                  EInstructionReportErrorCode = "INVALID_RUNNER"
	IREC_BET_TAKEN_OR_LAPSED             EInstructionReportErrorCode = "BET_TAKEN_OR_LAPSED"
	IREC_BET_IN_PROGRESS                 EInstructionReportErrorCode = "BET_IN_PROGRESS"
	IREC_RUNNER_REMOVED                  EInstructionReportErrorCode = "RUNNER_REMOVED"
	IREC_MARKET_NOT_OPEN_FOR_BETTING     EInstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"
	IREC_LOSS_LIMIT_EXCEEDED             EInstructionReportErrorCode = "LOSS_LIMIT_EXCEEDED"
	IREC_MARKET_NOT_OPEN_FOR_BSP_BETTING EInstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BSP_BETTING"
	IREC_INVALID_PRICE_EDIT              EInstructionReportErrorCode = "INVALID_PRICE_EDIT"
	IREC_INVALID_ODDS                    EInstructionReportErrorCode = "INVALID_ODDS"
	IREC_INSUFFICIENT_FUNDS              EInstructionReportErrorCode = "INSUFFICIENT_FUNDS"
	IREC_INVALID_PERSISTENCE_TYPE        EInstructionReportErrorCode = "INVALID_PERSISTENCE_TYPE"
	IREC_ERROR_IN_MATCHER                EInstructionReportErrorCode = "ERROR_IN_MATCHER"
	IREC_INVALID_BACK_LAY_COMBINATION    EInstructionReportErrorCode = "INVALID_BACK_LAY_COMBINATION"
	IREC_ERROR_IN_ORDER                  EInstructionReportErrorCode = "ERROR_IN_ORDER"
	IREC_INVALID_BID_TYPE                EInstructionReportErrorCode = "INVALID_BID_TYPE"
	IREC_INVALID_BET_ID                  EInstructionReportErrorCode = "INVALID_BET_ID"
	IREC_CANCELLED_NOT_PLACED            EInstructionReportErrorCode = "CANCELLED_NOT_PLACED"
	IREC_RELATED_ACTION_FAILED           EInstructionReportErrorCode = "RELATED_ACTION_FAILED"
	IREC_NO_ACTION_REQUIRED              EInstructionReportErrorCode = "NO_ACTION_REQUIRED"
	IREC_TIME_IN_FORCE_CONFLICT          EInstructionReportErrorCode = "TIME_IN_FORCE_CONFLICT"
	IREC_UNEXPECTED_PERSISTENCE_TYPE     EInstructionReportErrorCode = "UNEXPECTED_PERSISTENCE_TYPE"
	IREC_INVALID_ORDER_TYPE              EInstructionReportErrorCode = "INVALID_ORDER_TYPE"
	IREC_UNEXPECTED_MIN_FILL_SIZE        EInstructionReportErrorCode = "UNEXPECTED_MIN_FILL_SIZE"
	IREC_INVALID_CUSTOMER_ORDER_REF      EInstructionReportErrorCode = "INVALID_CUSTOMER_ORDER_REF"
	IREC_INVALID_MIN_FILL_SIZE           EInstructionReportErrorCode = "INVALID_MIN_FILL_SIZE"
	ErrUnknownInstructionReportStatus    EInstructionReportStatus    = "Unknown instructionReportStatus value"
)

type EItemClass string

const (
	IC_UNKNOWN EItemClass = "UNKNOWN"
)

type EMarketType string

const (
	MT_A              EMarketType = "A"
	MT_L              EMarketType = "L"
	MT_O              EMarketType = "O"
	MT_R              EMarketType = "R"
	MT_NOT_APPLICABLE EMarketType = "NOT_APPLICABLE"
)

type ELoginStatus string

const (
	LS_SUCCESS                                 ELoginStatus = "SUCCESS"
	LS_INVALID_USERNAME_OR_PASSWORD            ELoginStatus = "INVALID_USERNAME_OR_PASSWORD"
	LS_ACCOUNT_NOW_LOCKED                      ELoginStatus = "ACCOUNT_NOW_LOCKED"
	LS_ACCOUNT_ALREADY_LOCKED                  ELoginStatus = "ACCOUNT_ALREADY_LOCKED"
	LS_PENDING_AUTH                            ELoginStatus = "PENDING_AUTH"
	LS_TELBET_TERMS_CONDITIONS_NA              ELoginStatus = "TELBET_TERMS_CONDITIONS_NA"
	LS_DUPLICATE_CARDS                         ELoginStatus = "DUPLICATE_CARDS"
	LS_SECURITY_QUESTION_WRONG_3X              ELoginStatus = "SECURITY_QUESTION_WRONG_3X"
	LS_KYC_SUSPEND                             ELoginStatus = "KYC_SUSPEND"
	LS_SUSPENDED                               ELoginStatus = "SUSPENDED"
	LS_CLOSED                                  ELoginStatus = "CLOSED"
	LS_SELF_EXCLUDED                           ELoginStatus = "SELF_EXCLUDED"
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK    ELoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_DK"
	LS_NOT_AUTHORIZED_BY_REGULATOR_DK          ELoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_DK"
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT    ELoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_IT"
	LS_NOT_AUTHORIZED_BY_REGULATOR_IT          ELoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_IT"
	LS_SECURITY_RESTRICTED_LOCATION            ELoginStatus = "SECURITY_RESTRICTED_LOCATION"
	LS_BETTING_RESTRICTED_LOCATION             ELoginStatus = "BETTING_RESTRICTED_LOCATION"
	LS_TRADING_MASTER                          ELoginStatus = "TRADING_MASTER"
	LS_TRADING_MASTER_SUSPENDED                ELoginStatus = "TRADING_MASTER_SUSPENDED"
	LS_AGENT_CLIENT_MASTER                     ELoginStatus = "AGENT_CLIENT_MASTER"
	LS_AGENT_CLIENT_MASTER_SUSPENDED           ELoginStatus = "AGENT_CLIENT_MASTER_SUSPENDED"
	LS_DANISH_AUTHORIZATION_REQUIRED           ELoginStatus = "DANISH_AUTHORIZATION_REQUIRED"
	LS_SPAIN_MIGRATION_REQUIRED                ELoginStatus = "SPAIN_MIGRATION_REQUIRED"
	LS_DENMARK_MIGRATION_REQUIRED              ELoginStatus = "DENMARK_MIGRATION_REQUIRED"
	LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED       ELoginStatus = "SPANISH_TERMS_ACCEPTANCE_REQUIRED"
	LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED    ELoginStatus = "ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED"
	LS_CERT_AUTH_REQUIRED                      ELoginStatus = "CERT_AUTH_REQUIRED"
	LS_CHANGE_PASSWORD_REQUIRED                ELoginStatus = "CHANGE_PASSWORD_REQUIRED"
	LS_PERSONAL_MESSAGE_REQUIRED               ELoginStatus = "PERSONAL_MESSAGE_REQUIRED"
	LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED ELoginStatus = "INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED"
	LS_EMAIL_LOGIN_NOT_ALLOWED                 ELoginStatus = "EMAIL_LOGIN_NOT_ALLOWED"
	LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL     ELoginStatus = "MULTIPLE_USERS_WITH_SAME_CREDENTIAL"
	LS_ACCOUNT_PENDING_PASSWORD_CHANGE         ELoginStatus = "ACCOUNT_PENDING_PASSWORD_CHANGE"
	LS_TEMPORARY_BAN_TOO_MANY_REQUESTS         ELoginStatus = "TEMPORARY_BAN_TOO_MANY_REQUESTS"
	ErrUnknownLoginStatus                      ELoginStatus = "Unknown loginStatus value"
)

type EMarketBettingType string
type EMarketBettingTypes []EMarketBettingType

const (
	MBT_ODDS                       EMarketBettingType = "ODDS"
	MBT_LINE                       EMarketBettingType = "LINE"
	MBT_ASIAN_HANDICAP_DOUBLE_LINE EMarketBettingType = "ASIAN_HANDICAP_DOUBLE_LINE"
	MBT_ASIAN_HANDICAP_SINGLE_LINE EMarketBettingType = "ASIAN_HANDICAP_SINGLE_LINE"
	MBT_FIXED_ODDS                 EMarketBettingType = "FIXED_ODDS"
	ErrUnknownMarketBettingType    EMarketBettingType = "Unknown marketBettingType value"
)

type EMarketProjection string

const (
	MP_COMPETITION             EMarketProjection = "COMPETITION"
	MP_EVENT                   EMarketProjection = "EVENT"
	MP_EVENT_TYPE              EMarketProjection = "EVENT_TYPE"
	MP_MARKET_START_TIME       EMarketProjection = "MARKET_START_TIME"
	MP_MARKET_DESCRIPTION      EMarketProjection = "MARKET_DESCRIPTION"
	MP_RUNNER_DESCRIPTION      EMarketProjection = "RUNNER_DESCRIPTION"
	MP_RUNNER_METADATA         EMarketProjection = "RUNNER_METADATA"
	ErrUnknownMarketProjection EMarketProjection = "Unknown marketProjection value"
)

type EMarketSort string

const (
	MS_MINIMUM_TRADED    EMarketSort = "MINIMUM_TRADED"
	MS_MAXIMUM_TRADED    EMarketSort = "MAXIMUM_TRADED"
	MS_MINIMUM_AVAILABLE EMarketSort = "MINIMUM_AVAILABLE"
	MS_MAXIMUM_AVAILABLE EMarketSort = "MAXIMUM_AVAILABLE"
	MS_FIRST_TO_START    EMarketSort = "FIRST_TO_START"
	MS_LAST_TO_START     EMarketSort = "LAST_TO_START"
	ErrUnknownMarketSort EMarketSort = "Unknown marketSort value"
)

type EMarketStatus string

const (
	MS_INACTIVE           EMarketStatus = "INACTIVE"
	MS_OPEN               EMarketStatus = "OPEN"
	MS_SUSPENDED          EMarketStatus = "SUSPENDED"
	MS_CLOSED             EMarketStatus = "CLOSED"
	ErrUnknowMarketStatus EMarketStatus = "Unknown marketStatus value"
)

type EMatchProjection string

const (
	MP_NO_ROLLUP              EMatchProjection = "NO_ROLLUP"
	MP_ROLLED_UP_BY_PRICE     EMatchProjection = "ROLLED_UP_BY_PRICE"
	MP_ROLLED_UP_BY_AVG_PRICE EMatchProjection = "ROLLED_UP_BY_AVG_PRICE"
	ErrUnknownMatchProjection EMatchProjection = "Unknown matchProjection value"
)

type EOrderBy string

const (
	OB_BY_BET          EOrderBy = "BY_BET"
	OB_BY_MARKET       EOrderBy = "BY_MARKET"
	OB_BY_MATCH_TIME   EOrderBy = "BY_MATCH_TIME"
	OB_BY_PLACE_TIME   EOrderBy = "BY_PLACE_TIME"
	OB_BY_SETTLED_TIME EOrderBy = "BY_SETTLED_TIME"
	OB_BY_VOID_TIME    EOrderBy = "BY_VOID_TIME"
	ErrUnknownOrderBy  EOrderBy = "Unknown orderBy value"
)

type EOrderProjection string

const (
	OP_ALL                    EOrderProjection = "ALL"
	OP_EXECUTABLE             EOrderProjection = "EXECUTABLE"
	OP_EXECUTION_COMPLETE     EOrderProjection = "EXECUTION_COMPLETE"
	ErrUnknownOrderProjection EOrderProjection = "Unknown orderProjection value"
)

type EOrderStatus string

const (
	OS_PENDING            EOrderStatus = "PENDING"
	OS_EXECUTION_COMPLETE EOrderStatus = "EXECUTION_COMPLETE"
	OS_EXECUTABLE         EOrderStatus = "EXECUTABLE"
	OS_EXPIRED            EOrderStatus = "EXPIRED"
	ErrOrderStatus        EOrderStatus = "Unknown orderStatus value"
)

type EPriceData string

const (
	PD_SP_AVAILABLE     EPriceData = "SP_AVAILABLE"
	PD_SP_TRADED        EPriceData = "SP_TRADED"
	PD_EX_BEST_OFFERS   EPriceData = "EX_BEST_OFFERS"
	PD_EX_ALL_OFFERS    EPriceData = "EX_ALL_OFFERS"
	PD_EX_TRADED        EPriceData = "EX_TRADED"
	ErrUnknownPriceData EPriceData = "Unknown priceData value"
)

type EPersistenceType string

const (
	PT_LAPSE                  EPersistenceType = "LAPSE"
	PT_PERSIST                EPersistenceType = "PERSIST"
	PT_MARKET_ON_CLOSE        EPersistenceType = "MARKET_ON_CLOSE"
	ErrUnknownPersistenceType EPersistenceType = "Unknown persistenceType value"
)

type ERollupModel string

const (
	RM_STAKE              ERollupModel = "STAKE"
	RM_PAYOUT             ERollupModel = "PAYOUT"
	RM_MANAGED_LIABILITY  ERollupModel = "MANAGED_LIABILITY"
	RM_NONE               ERollupModel = "NONE"
	ErrUnknownRollupModel ERollupModel = "Unknown rollupModel value"
)

type ERunnerStatus string

const (
	RS_ACTIVE              ERunnerStatus = "ACTIVE"
	RS_WINNER              ERunnerStatus = "WINNER"
	RS_LOSER               ERunnerStatus = "LOSER"
	RS_PLACED              ERunnerStatus = "PLACED"
	RS_REMOVED_VACANT      ERunnerStatus = "REMOVED_VACANT"
	RS_REMOVED             ERunnerStatus = "REMOVED"
	RS_HIDDEN              ERunnerStatus = "HIDDEN"
	ErrUnknownRunnerStatus ERunnerStatus = "Unknown runnerStatus value"
)

type ESide string

const (
	S_BACK         ESide = "BACK"
	S_LAY          ESide = "LAY"
	ErrUnknownSide ESide = "Unknown side value. Supported BACK or LAY only"
)

type ESortDir string

const (
	EARLIEST_TO_LATEST ESortDir = "EARLIEST_TO_LATEST"
	LATEST_TO_EARLIEST ESortDir = "LATEST_TO_EARLIEST"
)

type ETimeGranularity string

const (
	TG_DAYS                   ETimeGranularity = "DAYS"
	TG_HOURS                  ETimeGranularity = "HOURS"
	TG_MINUTES                ETimeGranularity = "MINUTES"
	ErrUnknownTimeGranularity ETimeGranularity = "Unknown timeGranularity value"
)

type ETimeInForce string

const (
	ETIF_FILL_OR_KILL ETimeInForce = "FILL_OR_KILL"
)

type EOrderType string

const (
	OT_LIMIT           EOrderType = "LIMIT"
	OT_LIMIT_ON_CLOSE  EOrderType = "LIMIT_ON_CLOSE"
	OT_MARKET_ON_CLOSE EOrderType = "MARKET_ON_CLOSE"
)

type EBetTargetType string

const (
	BTT_BACKERS_PROFIT EBetTargetType = "BACKERS_PROFIT"
	BTT_PAYOUT         EBetTargetType = "PAYOUT"
)
