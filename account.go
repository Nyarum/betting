package betting

type DeveloperAppVersion struct {
	Owner                string
	VersionID            int64
	Version              string
	ApplicationKey       string
	DelayData            bool
	SubscriptionRequired bool
	OwnerManaged         bool
	Active               bool
}

type DeveloperApp struct {
	AppName     string
	AppID       int64
	AppVersions []DeveloperAppVersion
}

func (b *Betting) GetAppKeys() ([]DeveloperApp, error) {
	return nil, nil
}
