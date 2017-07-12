package betting

type UpdateKey struct {
	EventID                     string `json:"eventId,omitempty"`
	LastUpdateSequenceProcessed int64  `json:"lastUpdateSequenceProcessed,omitempty"`
}

type ScoresParams struct {
	EventIDs     []string       `json:"eventIds,omitempty"`
	EventTypeIDs []string       `json:"eventTypeIds,omitempty"`
	EventStatus  []EEventStatus `json:"eventStatus,omitempty"`
	UpdateKeys   []UpdateKey    `json:"updateKeys,omitempty"`
}
