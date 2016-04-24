package betting

import "time"

type DataRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"from,omitempty"`
}

type Filter struct {
	Wallet        Wallet      `json:"wallet,omitempty"`
	Locale        string      `json:"wallet,omitempty"`
	FromRecord    int         `json:"fromRecord,omitempty"`
	RecordCount   int         `json:"recordCount,omitempty"`
	ItemDataRange DataRange   `json:"itemDateRange,omitempty"`
	IncludeItem   IncludeItem `json:"recordCount,omitempty"`
	FromCurrency  string      `json:"fromCurrency,omitempty"`
	From          Wallet      `json:"from,omitempty"`
	To            Wallet      `json:"to,omitempty"`
	Amount        float64     `json:"amount,omitempty"`
}
