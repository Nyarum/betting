package betting

type ScoresListMethod string

const (
	ListAvailableEvents ScoresListMethod = "listAvailableEvents"
	ListScores          ScoresListMethod = "listScores"
	ListIncidents       ScoresListMethod = "listIncidents"
)

func (s *Scores) ListAvailableEvents(params ScoresParams) (availableEvents []AvailableEvent, err error) {
	err = s.Request(&availableEvents, ListAvailableEvents, &params)

	return
}

func (s *Scores) ListScores(params ScoresParams) (score []Score, err error) {
	err = s.Request(&score, ListScores, &params)

	return
}

func (s *Scores) ListIncidents(params ScoresParams) (score []Incidents, err error) {
	err = s.Request(&score, ListIncidents, &params)

	return
}
