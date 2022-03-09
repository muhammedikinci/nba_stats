package simulation

type team struct {
	Name    string   `json:"name"`
	Players []player `json:"players"`
	Points  int      `json:"points"`
}

type player struct {
	Name          string   `json:"name"`
	Position      string   `json:"position"`
	TeamId        int      `json:"team_id"`
	PTS           float64  `json:"-"`
	REB           float64  `json:"-"`
	AST           float64  `json:"-"`
	STL           float64  `json:"-"`
	BLK           float64  `json:"-"`
	Actions       []string `json:"actions"`
	Points        int      `json:"points"`
	Blocks        int      `json:"blocks"`
	Assists       int      `json:"assists"`
	SuccessShoots int      `json:"successShoots"`
	FailShoots    int      `json:"failShoots"`
}
