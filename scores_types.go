package betting

import "time"

type UpdateContext struct {
	EventTime      string    `json:"eventTime,omitempty"`
	LastUpdated    time.Time `json:"lastUpdated,omitempty"`
	UpdateSequence int64     `json:"updateSequence,omitempty"`
	UpdateType     string    `json:"updateType,omitempty"`
}

type UpdateKey struct {
	EventID                     string `json:"eventId,omitempty"`
	LastUpdateSequenceProcessed int64  `json:"lastUpdateSequenceProcessed,omitempty"`
}

type Incident struct {
	UpdateContext UpdateContext     `json:"updateContext,omitempty"`
	Values        map[string]string `json:"values,omitempty"`
}

type Incidents struct {
	EventID      string             `json:"eventId,omitempty"`
	EventTypeID  string             `json:"eventTypeId,omitempty"`
	EventStatus  EEventStatus       `json:"eventStatus,omitempty"`
	ResponseCode EResponseCode      `json:"responseCode,omitempty"`
	Incidents    map[int64]Incident `json:"incidents,omitempty"`
}

type Score struct {
	EventID       string            `json:"eventId,omitempty"`
	EventTypeID   string            `json:"eventTypeId,omitempty"`
	EventStatus   EEventStatus      `json:"eventStatus,omitempty"`
	ResponseCode  EResponseCode     `json:"responseCode,omitempty"`
	UpdateContext UpdateContext     `json:"updateContext,omitempty"`
	Values        map[string]string `json:"values,omitempty"`
}

type AvailableEvent struct {
	EventID     string       `json:"eventId,omitempty"`
	EventTypeID string       `json:"eventTypeId,omitempty"`
	EventStatus EEventStatus `json:"eventStatus,omitempty"`
}
