package model

type League struct {
	IdLeague        byte   `json:"idLeague"`
	LeagueName      string `json:"leagueName"`
	Country         string `json:"country"`
	LogoLeagueUrl   string `json:"logoLeagueUrl"`
	LeaderStandings string `json:"leaderStandings"`
	Teams           []Club `json:"-"`
}
