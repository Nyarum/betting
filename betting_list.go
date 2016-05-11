package betting

import "time"

type Competition struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompetitionResult struct {
	Competition       Competition `json:"competition,omitempty"`
	MarketCount       int         `json:"marketCount,omitempty"`
	CompetitionRegion string      `json:"competitionRegion,omitempty"`
}

// ListCompetitions to get competitions associated with the markets selected by the filter
func (b *Betting) ListCompetitions(filter Filter) (competitionResults []CompetitionResult, err error) {
	err = b.Request(&competitionResults, BettingURL, "listCompetitions", &filter)
	if err != nil {
		return
	}

	return
}

type CountryCodeResult struct {
	CountryCode string `json:"countryCode,omitempty"`
	MarketCount int    `json:"marketCount,omitempty"`
}

// ListCountries to get a list of countries associated with the markets selected by the filter
func (b *Betting) ListCountries(filter Filter) (countryCodeResults []CountryCodeResult, err error) {
	err = b.Request(&countryCodeResults, BettingURL, "listCountries", &filter)
	if err != nil {
		return
	}

	return
}

type PriceSize struct {
	Price float64
	Size  float64
}

type CurrentOrderSummary struct {
	BetID               string          `json:"betId"`
	MarketID            string          `json:"marketId"`
	SelectionID         int64           `json:"selectionId"`
	Handicap            float64         `json:"handicap"`
	PriceSize           PriceSize       `json:"priceSize"`
	BspLiability        float64         `json:"bspLiability"`
	Side                Side            `json:"side"`
	Status              OrderStatus     `json:"status"`
	PersistenceType     PersistenceType `json:"persistenceType"`
	OrderType           OrderType       `json:"orderType"`
	PlacedDate          time.Time       `json:"placedDate"`
	MatchedDate         time.Time       `json:"matchedDate"`
	AveragePriceMatched float64         `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64         `json:"sizeMatched,omitempty"`
	SizeRemaining       float64         `json:"sizeRemaining,omitempty"`
	SizeLapsed          float64         `json:"sizeLapsed,omitempty"`
	SizeCancelled       float64         `json:"sizeCancelled,omitempty"`
	SizeVoided          float64         `json:"sizeVoided,omitempty"`
	RegulatorAuthCode   string          `json:"regulatorAuthCode,omitempty"`
	RegulatorCode       string          `json:"regulatorCode,omitempty"`
}

type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders,omitempty"`
	MoreAvailable bool                  `json:"moreAvailable,omitempty"`
}

// ListCurrentOrders to get a list of your current orders
func (b *Betting) ListCurrentOrders(filter Filter) (currentOrderSummaryReport CurrentOrderSummaryReport, err error) {
	err = b.Request(&currentOrderSummaryReport, BettingURL, "listCurrentOrders", &filter)
	if err != nil {
		return
	}

	return
}
