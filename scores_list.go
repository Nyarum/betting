package betting

import "time"

type ScoresListMethod string

const (
	ListAvailableEvents ScoresListMethod = "listAvailableEvents"
	ListScores          ScoresListMethod = "listScores"
	ListIncidents       ScoresListMethod = "listIncidents"
)

type AvailableEvent struct {
	EventID     string       `json:"eventId,omitempty"`
	EventTypeID string       `json:"eventTypeId,omitempty"`
	EventStatus EEventStatus `json:"eventStatus,omitempty"`
}

func (s *Scores) ListAvailableEvents(params ScoresParams) (availableEvents []AvailableEvent, err error) {
	err = s.Request(&availableEvents, ListAvailableEvents, &params)

	return
}

type UpdateContext struct {
	EventTime      string    `json:"eventTime,omitempty"`
	LastUpdated    time.Time `json:"lastUpdated,omitempty"`
	UpdateSequence int64     `json:"updateSequence,omitempty"`
	UpdateType     string    `json:"updateType,omitempty"`
}

type Score struct {
	EventID       string            `json:"eventId,omitempty"`
	EventTypeID   string            `json:"eventTypeId,omitempty"`
	EventStatus   EEventStatus      `json:"eventStatus,omitempty"`
	ResponseCode  EResponseCode     `json:"responseCode,omitempty"`
	UpdateContext UpdateContext     `json:"updateContext,omitempty"`
	Values        map[string]string `json:"values,omitempty"`
}

func (s *Scores) ListScores(params ScoresParams) (score []Score, err error) {
	err = s.Request(&score, ListScores, &params)

	return
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

func (s *Scores) ListIncidents(params ScoresParams) (score []Incidents, err error) {
	err = s.Request(&score, ListIncidents, &params)

	return
}
