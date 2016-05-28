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

type ItemDescription struct {
	EventTypeDesc   string    `json:"eventTypeDesc,omitempty"`
	EventDesc       string    `json:"eventDesc,omitempty"`
	MarketDesc      string    `json:"marketDesc,omitempty"`
	MarketType      string    `json:"marketType,omitempty"`
	MarketStartTime time.Time `json:"marketStartTime,omitempty"`
	RunnerDesc      string    `json:"runnerDesc,omitempty"`
	NumberOfWinners int       `json:"numberOfWinners,omitempty"`
	EachWayDivisor  float64   `json:"eachWayDivisor,omitempty"`
}

type ClearedOrderSummary struct {
	EventTypeID     string          `json:"eventTypeId,omitempty"`
	EventID         string          `json:"eventId,omitempty"`
	MarketID        string          `json:"marketId,omitempty"`
	SelectionID     int64           `json:"selectionId,omitempty"`
	Handicap        float64         `json:"handicap,omitempty"`
	BetID           string          `json:"betId,omitempty"`
	PlacedDate      time.Time       `json:"placedDate,omitempty"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
	OrderType       OrderType       `json:"orderType,omitempty"`
	Side            Side            `json:"side,omitempty"`
	ItemDescription ItemDescription `json:"itemDescription,omitempty"`
	BetOutcome      string          `json:"betOutcome,omitempty"`
	PriceRequested  float64         `json:"priceRequested,omitempty"`
	SettledDate     time.Time       `json:"settledDate,omitempty"`
	LastMatchedDate time.Time       `json:"lastMatchedDate,omitempty"`
	BetCount        int             `json:"betCount,omitempty"`
	Commission      float64         `json:"commission,omitempty"`
	PriceMatched    float64         `json:"priceMatched,omitempty"`
	PriceReduced    bool            `json:"priceReduced,omitempty"`
	SizeSettled     float64         `json:"sizeSettled,omitempty"`
	Profit          float64         `json:"profit,omitempty"`
	SizeCancelled   float64         `json:"sizeCancelled,omitempty"`
}

type ClearedOrderSummaryReport struct {
	ClearedOrders []ClearedOrderSummary `json:"clearedOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

// ListClearedOrders to get a list of settled bets based on the bet status, ordered by settled date
func (b *Betting) ListClearedOrders(filter Filter) (clearedOrderSummaryReport ClearedOrderSummaryReport, err error) {
	err = b.Request(&clearedOrderSummaryReport, BettingURL, "listClearedOrders", &filter)
	if err != nil {
		return
	}

	return
}

type Event struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	CountryCode string    `json:"countryCode,omitempty"`
	Timezone    string    `json:"timezone,omitempty"`
	Venue       string    `json:"venue,omitempty"`
	OpenDate    time.Time `json:"openDate,omitempty"`
}

type EventResult struct {
	Event       Event `json:"event,omitempty"`
	MarketCount int   `json:"marketCount,omitempty"`
}

// ListEvents to get a list of all events
func (b *Betting) ListEvents(filter Filter) (eventResult []EventResult, err error) {
	err = b.Request(&eventResult, BettingURL, "listEvents", &filter)
	if err != nil {
		return
	}

	return
}

type EventType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EventTypeResult struct {
	EventType   EventType `json:"eventType,omitempty"`
	MarketCount int       `json:"marketCount,omitempty"`
}

// ListEventTypes to get a list of all event types
func (b *Betting) ListEventTypes(filter Filter) (eventTypeResult []EventTypeResult, err error) {
	err = b.Request(&eventTypeResult, BettingURL, "listEventTypes", &filter)
	if err != nil {
		return
	}

	return
}
