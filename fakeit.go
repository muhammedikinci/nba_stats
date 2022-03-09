package main

import (
	"math/rand"
	"nba_stats/model"
	"nba_stats/repository"
	"strings"
)

var playerNames = []string{"Joel Embiid", "LeBron James", "Giannis Antetokounmpo", "DeMar DeRozan", "Trae Young", "Rudy Gobert", "Nikola Jokic", "Domantas Sabonis", "Clint Capela", "Giannis Antetokounmpo", "Chris Paul", "James Harden", "Dejounte Murray", "Trae Young", "Luka Doncic", "Robert Williams III", "Rudy Gobert", "Jaren Jackson Jr.", "Jakob Poeltl", "Mo Bamba", "Dejounte Murray", "Chris Paul", "Gary Trent Jr.", "Tyrese Haliburton", "Matisse Thybulle", "Rudy Gobert", "Jarrett Allen", "Deandre Ayton", "Montrezl Harrell", "Jakob Poeltl", "Stephen Curry", "Buddy Hield", "Fred VanVleet", "Patty Mills", "Malik Beasley", "Luke Kennard", "P.J. Tucker", "Cameron Johnson", "Grant Williams", "Mike Muscala", "Nikola Jokic", "Giannis Antetokounmpo", "Joel Embiid", "Luka Doncic", "LeBron James", "Robert Williams III", "Mitchell Robinson", "Rudy Gobert", "Dwight Powell", "Jaxson Hayes", "Joel Embiid", "Luka Doncic", "Giannis Antetokounmpo", "Trae Young", "Ja Morant", "Steven Adams", "Andre Drummond", "Mitchell Robinson", "Hassan Whiteside", "JaVale McGee", "LeBron James", "Ja Morant", "Giannis Antetokounmpo", "Devin Booker", "Josh Hart", "Rudy Gobert", "Cameron Oliver", "Mitchell Robinson", "Jonas Valanciunas", "Nikola Jokic", "Ja Morant", "Giannis Antetokounmpo", "Nikola Jokic", "Anthony Davis", "LeBron James", "Malik Beasley", "Patty Mills", "Devonte' Graham", "Buddy Hield", "Evan Fournier", "Clint Capela", "Jakob Poeltl", "Jarrett Allen", "Deandre Ayton", "Rudy Gobert", "LaMarcus Aldridge", "DeMar DeRozan", "Chris Paul", "Brandon Ingram", "Seth Curry", "Ryan Arcidiacono", "Zach Collins", "Gabriel Deck", "Enes Freedom", "Jared Harper", "Mikal Bridges", "Gary Payton II", "Daniel Gafford", "Jaxson Hayes", "Hassan Whiteside", "LaMarcus Aldridge", "Seth Curry", "Trae Young", "Chris Paul", "Darius Garland", "Joel Embiid", "Giannis Antetokounmpo", "Luka Doncic", "Nikola Jokic", "Shai Gilgeous-Alexander", "Nikola Jokic", "Khris Middleton", "Malcolm Brogdon", "Jerami Grant", "Norman Powell", "Fred VanVleet", "Christian Wood", "Luguentz Dort", "Seth Curry", "Jae'Sean Tate", "Joel Embiid", "Nikola Jokic", "Karl-Anthony Towns", "Julius Randle", "Domantas Sabonis", "Rudy Gobert", "Nikola Jokic", "Domantas Sabonis", "Clint Capela", "Jonas Valanciunas", "Nikola Jokic", "Julius Randle", "Domantas Sabonis", "Joel Embiid", "Karl-Anthony Towns", "Joel Embiid", "LeBron James", "Giannis Antetokounmpo", "DeMar DeRozan", "Luka Doncic", "Domantas Sabonis", "Giannis Antetokounmpo", "Joel Embiid", "Wendell Carter Jr.", "Christian Wood", "Luka Doncic", "LeBron James", "Giannis Antetokounmpo", "Jimmy Butler", "Brandon Ingram", "DeMar DeRozan", "Trae Young", "Luka Doncic", "Ja Morant", "Jayson Tatum", "Luka Doncic", "Dejounte Murray", "Jayson Tatum", "James Harden", "Josh Giddey", "Chris Paul", "James Harden", "Dejounte Murray", "Trae Young", "Luka Doncic", "LeBron James", "Karl-Anthony Towns", "Anthony Edwards", "Andrew Wiggins", "Cade Cunningham", "Karl-Anthony Towns", "LeBron James", "Cade Cunningham", "Dwight Howard", "Anthony Edwards", "LeBron James", "Cade Cunningham", "Karl-Anthony Towns", "Anthony Edwards", "Andrew Wiggins", "Scottie Barnes", "Evan Mobley", "Cade Cunningham", "Franz Wagner", "Josh Giddey", "Cade Cunningham", "Franz Wagner", "Jalen Green", "Scottie Barnes", "Evan Mobley", "Josh Giddey", "Evan Mobley", "Scottie Barnes", "Omer Yurtseven", "Cade Cunningham", "LeBron James", "Carmelo Anthony", "Kevin Durant", "Russell Westbrook", "James Harden", "Dwight Howard", "LeBron James", "DeAndre Jordan", "LaMarcus Aldridge", "Kevin Love", "Chris Paul", "LeBron James", "Russell Westbrook", "Rajon Rondo", "Kyle Lowry"}
var positions = []string{"PG", "SG", "SF", "PF", "C"}

var teams = []model.Team{
	{Name: "76ers"},
	{Name: "Bucks"},
	{Name: "Bulls"},
	{Name: "Cavaliers"},
	{Name: "Celtics"},
	{Name: "Clippers"},
}

func generateData() {
	repository := repository.NewRepository()

	repository.CreatePlayersTable()
	repository.CreateTeamsTable()

	repository.RemovePlayer()
	repository.RemoveTeams()

	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

	for _, team := range teams {
		repository.SaveTeam(team)
	}

	getTeams, err := repository.GetTeams()

	if err != nil {
		panic(err)
	}

	for _, team := range getTeams {
		for _, position := range positions {
			randomName := rand.Intn(len(playerNames))
			randomSurname := rand.Intn(len(playerNames))

			fullName := strings.Split(playerNames[randomName], " ")[0]
			fullName += " " + strings.Split(playerNames[randomSurname], " ")[1]

			player := model.Player{
				Name:     fullName,
				Position: position,
				TeamId:   team.Id,
				PTS:      rand.Float64(),
				REB:      rand.Float64(),
				AST:      rand.Float64(),
				STL:      rand.Float64(),
				BLK:      rand.Float64(),
			}

			repository.SavePlayer(player)
		}
	}
}
