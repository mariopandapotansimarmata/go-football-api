package data

import "github.com/mariopandapotansimarmata/go-api-responsi-mobile/model"

var ListLeague = []model.League{
	{
		IdLeague:        1,
		LeagueName:      "Premier League",
		Country:         "England",
		LogoLeagueUrl:   "premier_league_logo.png",
		LeaderStandings: "premier_league_standings",
		Teams:           ListPremierLeagueTeams,
	},
	{
		IdLeague:        2,
		LeagueName:      "La Liga",
		Country:         "Spain",
		LogoLeagueUrl:   "la_liga_logo.png",
		LeaderStandings: "la_liga_standings",
		Teams:           ListLaLigaTeams,
	},
	{
		IdLeague:        3,
		LeagueName:      "Bundesliga",
		Country:         "Germany",
		LogoLeagueUrl:   "bundesliga_logo.png",
		LeaderStandings: "bundesliga_standings",
		Teams:           ListBundeligaTeams,
	},
	{
		IdLeague:        4,
		LeagueName:      "Serie A",
		Country:         "Italy",
		LogoLeagueUrl:   "serie_a_logo.png",
		LeaderStandings: "serie_a_standings",
		Teams:           ListSerieATeams,
	},
	{
		IdLeague:        5,
		LeagueName:      "Ligue 1",
		Country:         "France",
		LogoLeagueUrl:   "ligue_1_logo.png",
		LeaderStandings: "ligue_1_standings",
		Teams:           ListLeague1Teams,
	},
	{
		IdLeague:        6,
		LeagueName:      "Eredivisie",
		Country:         "Netherlands",
		LogoLeagueUrl:   "eredivisie_logo.png",
		LeaderStandings: "eredivisie_standings",
		// Teams:           ListErediviseTeams,
	},
}
