package betting

import "time"

type DateRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"from,omitempty"`
}

type RunnerID struct {
	MarketID    string
	SelectionID int64
	Handicap    float64
}

type Filter struct {
	Wallet                       eWalletInternal          `json:"wallet,omitempty"`
	Locale                       string                   `json:"locale,omitempty"`
	FromRecord                   int                      `json:"fromRecord,omitempty"`
	RecordCount                  int                      `json:"recordCount,omitempty"`
	ItemDateRange                *DateRange               `json:"itemDateRange,omitempty"`
	IncludeItem                  eIncludeItemInternal     `json:"recordCount,omitempty"`
	FromCurrency                 string                   `json:"fromCurrency,omitempty"`
	From                         eWalletInternal          `json:"from,omitempty"`
	To                           eWalletInternal          `json:"to,omitempty"`
	Amount                       float64                  `json:"amount,omitempty"`
	BetIDs                       []string                 `json:"betIds,omitempty"`
	MarketIDs                    []string                 `json:"betIds,omitempty"`
	OrderProjection              eOrderProjectionInternal `json:"orderProjection,omitempty"`
	DateRange                    *DateRange               `json:"dateRange,omitempty"`
	OrderBy                      eOrderByInternal         `json:"orderBy,omitempty"`
	SortDir                      eSortDirInternal         `json:"sortDir,omitempty"`
	MarketFilter                 *MarketFilter            `json:"filter,omitempty"`
	BetStatus                    eBetStatusInternal       `json:"betStatus,omitempty"`
	EventTypeIDs                 []string                 `json:"eventTypeIds,omitempty"`
	EventIDs                     []string                 `json:"eventIds,omitempty"`
	RunnerIDs                    []RunnerID               `json:"runnerIds,omitempty"`
	Side                         eSideInternal            `json:"side,omitempty"`
	SettledDateRange             *DateRange               `json:"settledDateRange,omitempty"`
	GroupBy                      eGroupByInternal         `json:"groupBy,omitempty"`
	IncludeItemDescription       bool                     `json:"includeItemDescription,omitempty"`
	MaxResults                   int                      `json:"maxResults,omitempty"`
	IncludeSettledBets           bool                     `json:"includeSettledBets,omitempty"`
	MarketProfitAndLossMarketIds []string                 `json:"marketIds,omitempty"`
	TimeGranularity              eTimeGranularityInternal `json:"granularity,omitempty"`
}
