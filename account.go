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

type DeveloperAppKeysResult struct {
	Jsonrpc string
	Result  []DeveloperAppKey
}

// GetAppKeys for getting all developer keys from account
func (b *Betting) GetAppKeys() (developAppKeys []DeveloperAppKey, err error) {
	developAppKeysResult := &DeveloperAppKeysResult{}
	err = b.Request(developAppKeysResult, AccountURL, "AccountAPING/v1.0/getDeveloperAppKeys")
	if err != nil {
		return
	}

	developAppKeys = developAppKeysResult.Result

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

type AccountDetailsResult struct {
	Jsonrpc string
	Result  AccountDetails
}

// GetAccountDetails like get account details :)
func (b *Betting) GetAccountDetails() (accountDetails *AccountDetails, err error) {
	accountDetailsResult := &AccountDetailsResult{}
	err = b.Request(accountDetailsResult, AccountURL, "AccountAPING/v1.0/getAccountDetails")
	if err != nil {
		return
	}

	accountDetails = &accountDetailsResult.Result

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

type AccountFundsResult struct {
	Jsonrpc string
	Result  AccountFunds
}

// GetAccountFunds for getting balances of account
func (b *Betting) GetAccountFunds() (accountFunds *AccountFunds, err error) {
	accountFundsResult := &AccountFundsResult{}
	err = b.Request(accountFundsResult, AccountURL, "AccountAPING/v1.0/getAccountFunds")
	if err != nil {
		return
	}

	accountFunds = &accountFundsResult.Result

	return
}
