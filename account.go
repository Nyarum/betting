package betting

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
	err = b.Request(&developAppKeys, AccountURL, "createDeveloperAppKeys", nil)
	if err != nil {
		return
	}

	return
}

// GetAppKeys for getting all developer keys from account
func (b *Betting) GetAppKeys() (developAppKeys []DeveloperAppKey, err error) {
	err = b.Request(&developAppKeys, AccountURL, "getDeveloperAppKeys", nil)
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
	err = b.Request(&accountDetails, AccountURL, "getAccountDetails", nil)
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
	DiscountRate          *float64 `json:"omitempty"`
	PointsBalance         *int     `json:"omitempty"`
}

// GetAccountFunds for getting balances of account
func (b *Betting) GetAccountFunds(filter Filter) (accountFunds AccountFunds, err error) {
	err = b.Request(&accountFunds, AccountURL, "getAccountFunds", &filter)
	if err != nil {
		return
	}

	return
}
