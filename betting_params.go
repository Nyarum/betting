package betting

import "time"

type DateRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type RunnerID struct {
	MarketID    string
	SelectionID int64
	Handicap    float64
}

type EXBestOffersOverrides struct {
	BestPricesDepth          int          `json:"bestPricesDepth,omitempty"`
	RollupModel              ERollupModel `json:"rollupModel,omitempty"`
	RollupLimit              int          `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float64      `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    int          `json:"rollupLiabilityFactor,omitempty"`
}

type PriceProjection struct {
	PriceData             []EPriceData           `json:"priceData,omitempty"`
	EXBestOffersOverrides *EXBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                   `json:"virtualise,omitempty"`
	RolloverStakes        bool                   `json:"rolloverStakes,omitempty"`
}

type Filter struct {
	Wallet                       EWallet              `json:"wallet,omitempty"`
	Locale                       string               `json:"locale,omitempty"`
	FromRecord                   int                  `json:"fromRecord,omitempty"`
	RecordCount                  int                  `json:"recordCount,omitempty"`
	ItemDateRange                *DateRange           `json:"itemDateRange,omitempty"`
	IncludeItem                  EIncludeItem         `json:"recordCount,omitempty"`
	FromCurrency                 string               `json:"fromCurrency,omitempty"`
	From                         EWallet              `json:"from,omitempty"`
	To                           EWallet              `json:"to,omitempty"`
	Amount                       float64              `json:"amount,omitempty"`
	BetIDs                       []string             `json:"betIds,omitempty"`
	MarketIDs                    []string             `json:"marketIds,omitempty"`
	PriceProjection              *PriceProjection     `json:"priceProjection,omitempty"`
	OrderProjection              EOrderProjection     `json:"orderProjection,omitempty"`
	MarketProjection             *[]EMarketProjection `json:"marketProjection,omitempty"`
	DateRange                    *DateRange           `json:"dateRange,omitempty"`
	OrderBy                      EOrderBy             `json:"orderBy,omitempty"`
	SortDir                      ESortDir             `json:"sortDir,omitempty"`
	Sort                         EMarketSort          `json:"sort,omitempty"`
	MarketFilter                 *MarketFilter        `json:"filter,omitempty"`
	BetStatus                    EBetStatus           `json:"betStatus,omitempty"`
	EventTypeIDs                 []string             `json:"eventTypeIds,omitempty"`
	EventIDs                     []string             `json:"eventIds,omitempty"`
	RunnerIDs                    []RunnerID           `json:"runnerIds,omitempty"`
	Side                         ESide                `json:"side,omitempty"`
	SettledDateRange             *DateRange           `json:"settledDateRange,omitempty"`
	GroupBy                      EGroupBy             `json:"groupBy,omitempty"`
	IncludeItemDescription       bool                 `json:"includeItemDescription,omitempty"`
	MaxResults                   int                  `json:"maxResults,omitempty"`
	IncludeSettledBets           bool                 `json:"includeSettledBets,omitempty"`
	TimeGranularity              ETimeGranularity     `json:"granularity,omitempty"`
	PlaceOrdersMarketID          string               `json:"marketId,omitempty"`
	PlaceOrdersInstructions      []PlaceInstruction   `json:"instructions,omitempty"`
	CustomerOrderRefs            []string             `json:"customerOrderRefs,omitempty"`
	CustomerStrategyRefs         []string             `json:"customerStrategyRefs,omitempty"`
}
