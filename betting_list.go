package betting

type Competition struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompetitionResult struct {
	Competition       Competition `json:"competition,omitempty"`
	MarketCount       int         `json:"marketCount,omitempty"`
	CompetitionRegion string      `json:"competitionRegion,omitempty"`
}

// ListCompetitions to get competitions associated with the markets selected by filter
func (b *Betting) ListCompetitions(filter Filter) (competitionResult []CompetitionResult, err error) {
	err = b.Request(&competitionResult, BettingURL, "listCompetitions", &filter)
	if err != nil {
		return
	}

	return
}
