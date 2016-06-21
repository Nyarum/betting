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
	Wallet                 Wallet          `json:"wallet,omitempty"`
	Locale                 string          `json:"locale,omitempty"`
	FromRecord             int             `json:"fromRecord,omitempty"`
	RecordCount            int             `json:"recordCount,omitempty"`
	ItemDateRange          DateRange       `json:"itemDateRange,omitempty"`
	IncludeItem            IncludeItem     `json:"recordCount,omitempty"`
	FromCurrency           string          `json:"fromCurrency,omitempty"`
	From                   Wallet          `json:"from,omitempty"`
	To                     Wallet          `json:"to,omitempty"`
	Amount                 float64         `json:"amount,omitempty"`
	BetIDs                 []string        `json:"betIds,omitempty"`
	MarketIDs              []string        `json:"betIds,omitempty"`
	OrderProjection        OrderProjection `json:"orderProjection,omitempty"`
	DateRange              DateRange       `json:"dateRange,omitempty"`
	OrderBy                OrderBy         `json:"orderBy,omitempty"`
	SortDir                SortDir         `json:"sortDir,omitempty"`
	MarketFilter           MarketFilter    `json:"filter,omitempty"`
	BetStatus              BetStatus       `json:"betStatus,omitempty"`
	EventTypeIDs           []string        `json:"eventTypeIds,omitempty"`
	EventIDs               []string        `json:"eventIds,omitempty"`
	RunnerIDs              []RunnerID      `json:"runnerIds,omitempty"`
	Side                   Side            `json:"side,omitempty"`
	SettledDateRange       DateRange       `json:"settledDateRange,omitempty"`
	GroupBy                GroupBy         `json:"groupBy,omitempty"`
	IncludeItemDescription bool            `json:"includeItemDescription,omitempty"`
	MaxResults             int             `json:"maxResults,omitempty"`
}
