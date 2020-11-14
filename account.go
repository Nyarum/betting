package betting

import "time"

type DeveloperAppVersion struct {
	Owner                string
	VersionID            int
	Version              string
	ApplicationKey       string
	DelayData            bool
	SubscriptionRequired bool
	OwnerManaged         bool
	Active               bool
}

type DeveloperAppKey struct {
	AppName     string
	AppID       int64
	AppVersions []DeveloperAppVersion
}

// CreateAppKeys for create new developer app keys in account
func (b *Betting) CreateAppKeys() (developAppKeys []DeveloperAppKey, err error) {
	err = b.Request(&developAppKeys, b.AccountURL, "createDeveloperAppKeys", nil)
	if err != nil {
		return
	}

	return
}

type AccountDetails struct {
	CurrencyCode  string
	FirstName     string
	LastName      string
	LocaleCode    string
	Region        string
	Timezone      string
	DiscountRate  float64
	PointsBalance int
	CountryCode   string
}

// GetAccountDetails like get account details :)
func (b *Betting) GetAccountDetails() (accountDetails AccountDetails, err error) {
	err = b.Request(&accountDetails, b.AccountURL, "getAccountDetails", nil)
	if err != nil {
		return
	}

	return
}

type AccountFunds struct {
	AvailableToBetBalance float64
	Exposure              float64
	RetainedCommission    float64
	ExposureLimit         float64
	DiscountRate          float64 `json:"omitempty"`
	PointsBalance         int     `json:"omitempty"`
}

// GetAccountFunds for getting balances of account
func (b *Betting) GetAccountFunds(filter Filter) (accountFunds AccountFunds, err error) {
	err = b.Request(&accountFunds, b.AccountURL, "getAccountFunds", &filter)
	if err != nil {
		return
	}

	return
}

// GetAppKeys for getting all developer keys from account
func (b *Betting) GetAppKeys() (developAppKeys []DeveloperAppKey, err error) {
	err = b.Request(&developAppKeys, b.AccountURL, "getDeveloperAppKeys", nil)
	if err != nil {
		return
	}

	return
}

type StatementLegacyData struct {
	AvgPrice        float64     `json:"omitempty"`
	BetSize         float64     `json:"omitempty"`
	BetType         string      `json:"omitempty"`
	BetCategoryType string      `json:"omitempty"`
	CommissionRate  string      `json:"omitempty"`
	EventID         int64       `json:"omitempty"`
	EventTypeID     int64       `json:"omitempty"`
	FullMarketName  string      `json:"omitempty"`
	GrossBetAmount  float64     `json:"omitempty"`
	MarketName      string      `json:"omitempty"`
	MarketType      EMarketType `json:"omitempty"`
	PlacedDate      time.Time   `json:"omitempty"`
	SelectionID     int64       `json:"omitempty"`
	SelectionName   string      `json:"omitempty"`
	StartDate       time.Time   `json:"omitempty"`
	TransactionType string      `json:"omitempty"`
	TransactionID   int64       `json:"omitempty"`
	WinLose         string      `json:"omitempty"`
}

type StatementItem struct {
	RefID         string `json:"omitempty"`
	ItemDate      time.Time
	Amount        float64             `json:"omitempty"`
	Balance       float64             `json:"omitempty"`
	ItemClass     EItemClass          `json:"omitempty"`
	ItemClassData map[string]string   `json:"omitempty"`
	LegacyData    StatementLegacyData `json:"omitempty"`
}

type AccountStatementReport struct {
	AccountStatement []StatementItem
	MoreAvailable    bool
}

// GetAccountStatement about get account statements for 90 days
func (b *Betting) GetAccountStatement(filter Filter) (accountStatementReport AccountStatementReport, err error) {
	err = b.Request(&accountStatementReport, b.AccountURL, "getAccountStatement", &filter)
	if err != nil {
		return
	}

	return
}

type CurrencyRate struct {
	CurrencyCode string  `json:"omitempty"`
	Rate         float64 `json:"omitempty"`
}

// GetListCurrencyRates for get currency rates for hour
func (b *Betting) GetListCurrencyRates(filter Filter) (currencyRate []CurrencyRate, err error) {
	err = b.Request(&currencyRate, b.AccountURL, "listCurrencyRates", &filter)
	if err != nil {
		return
	}

	return
}

type TransferResponse struct {
	TransactionID string
}

// GetTransferFunds for transfer funds between UK and AUS wallets
func (b *Betting) GetTransferFunds(filter Filter) (transferResponse TransferResponse, err error) {
	err = b.Request(&transferResponse, b.AccountURL, "transferFunds", &filter)
	if err != nil {
		return
	}

	return
}
