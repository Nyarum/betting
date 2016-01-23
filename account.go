package betting

type DeveloperApp struct {
	Jsonrpc string
	Result  []struct {
		AppName     string
		AppID       int64
		AppVersions []struct {
			Owner                string
			VersionID            int64
			Version              string
			ApplicationKey       string
			DelayData            bool
			SubscriptionRequired bool
			OwnerManaged         bool
			Active               bool
		}
	}
}

func (b *Betting) GetAppKeys() (*DeveloperApp, error) {
	developApp := &DeveloperApp{}
	err := b.Request(developApp, AccountURL, "AccountAPING/v1.0/getDeveloperAppKeys")

	return developApp, err
}
