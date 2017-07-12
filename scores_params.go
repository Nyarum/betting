package betting

type ScoresParams struct {
	EventIDs     []string       `json:"eventIds,omitempty"`
	EventTypeIDs []string       `json:"eventTypeIds,omitempty"`
	EventStatus  []EEventStatus `json:"eventStatus,omitempty"`
	UpdateKeys   []UpdateKey    `json:"updateKeys,omitempty"`
}
