package data

import "github.com/mariopandapotansimarmata/go-api-responsi-mobile/model"

var ListLeague = []model.League{
	{
		IdLeague:        1,
		LeagueName:      "Premier League",
		Country:         "England",
		LogoLeagueUrl:   "https://b.fssta.com/uploads/application/soccer/competition-logos/EnglishPremierLeague.png",
		LeaderStandings: "Manchester City",
		Teams:           ListPremierLeagueTeams,
	},
	{
		IdLeague:        2,
		LeagueName:      "La Liga",
		Country:         "Spain",
		LogoLeagueUrl:   "https://upload.wikimedia.org/wikipedia/commons/thumb/0/0f/LaLiga_logo_2023.svg/2048px-LaLiga_logo_2023.svg.png",
		LeaderStandings: "Real Madrid",
		Teams:           ListLaLigaTeams,
	},
	{
		IdLeague:        3,
		LeagueName:      "Bundesliga",
		Country:         "Germany",
		LogoLeagueUrl:   "https://upload.wikimedia.org/wikinews/en/thumb/1/15/Bundesliga_logo.svg/544px-Bundesliga_logo.svg.png",
		LeaderStandings: "Bayern Munich",
		Teams:           ListBundeligaTeams,
	},
	{
		IdLeague:        4,
		LeagueName:      "Serie A",
		Country:         "Italy",
		LogoLeagueUrl:   "https://www.fifplay.com/img/public/serie-a-logo-transparent.png",
		LeaderStandings: "Juventus",
		Teams:           ListSerieATeams,
	},
	{
		IdLeague:        5,
		LeagueName:      "Ligue 1",
		Country:         "France",
		LogoLeagueUrl:   "https://upload.wikimedia.org/wikipedia/commons/f/fb/Ligue1_logo.png",
		LeaderStandings: "Paris Saint-Germain",
		Teams:           ListLeague1Teams,
	},
	{
		IdLeague:        6,
		LeagueName:      "Eredivisie",
		Country:         "Netherlands",
		LogoLeagueUrl:   "https://d3et0fncpz2hhr.cloudfront.net/Eredivisie-logo-blue-PNG.png",
		LeaderStandings: "Ajax",
		Teams:           ListErediviseTeams,
	},
}
