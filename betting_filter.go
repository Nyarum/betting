package betting

type MarketFilter struct {
	TextQuery          string                `json:"textQuery,omitempty"`
	ExchangeIDs        []string              `json:"exchangeIds,omitempty"`
	EventTypeIDs       []string              `json:"eventTypeIds,omitempty"`
	EventIDs           []string              `json:"eventIds,omitempty"`
	CompetitionIDs     []string              `json:"competitionIds,omitempty"`
	MarketIDs          []string              `json:"marketIds,omitempty"`
	Venues             []string              `json:"venues,omitempty"`
	BspOnly            *bool                 `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  *bool                 `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         *bool                 `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []EMarketBettingTypes `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string              `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string              `json:"marketTypeCodes,omitempty"`
	MarketStartTime    *DateRange            `json:"marketStartTime,omitempty"`
	WithOrders         []EOrderStatus        `json:"withOrders,omitempty"`
}
