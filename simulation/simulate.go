package simulation

import (
	"fmt"
	"math/rand"
	"nba_stats/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetTeams() ([]model.Team, error)
	GetPlayers(int) ([]model.Player, error)
}

type simulation struct {
	Repository   Repository
	IsSimulating bool
	Teams        []team
	Rounds       []round
	RuningRound  int
	GameTime     gameTimes
	Timer        *time.Ticker
}

type round struct {
	Matches []match `json:"matches"`
}

type gameTimes struct {
	oneMinute        int
	totalPassMinutes int
	secondCounter    int
}

func NewSimulation(repository Repository) *simulation {
	return &simulation{
		Repository: repository,
	}
}

func (sm *simulation) GetCurrentRound(c echo.Context) error {
	return c.JSON(http.StatusOK, sm.Rounds[sm.RuningRound])
}

func (sm *simulation) GetAllRounds(c echo.Context) error {
	return c.JSON(http.StatusOK, sm.Rounds)
}

func (sm *simulation) GetState(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":       sm.IsSimulating,
		"currentRound": sm.RuningRound,
	})
}

func (sm *simulation) StopSimulation(c echo.Context) error {
	sm.IsSimulating = false
	sm.Timer.Stop()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "simulation stopped",
	})
}

func (sm *simulation) StartSimulate(c echo.Context) error {
	if sm.IsSimulating {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  false,
			"message": "simulation already started",
		})
	}

	sm.IsSimulating = true

	err := sm.getDatas()

	if err != nil {
		panic(err)
	}

	runAll := false
	fmt.Println(c.QueryParam("runAll"))
	if c.QueryParam("runAll") == "true" {
		runAll = true
	}

	sm.createRounds()
	sm.simulate(runAll)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "simulation started",
	})
}

func (sm *simulation) getDatas() error {
	teams, err := sm.Repository.GetTeams()

	if err != nil {
		return err
	}

	sm.Teams = make([]team, len(teams))

	for i, team := range teams {
		players, err := sm.Repository.GetPlayers(team.Id)

		if err != nil {
			return err
		}

		convPlayers := []player{}

		for _, v := range players {
			convPlayers = append(convPlayers, player{
				Name:     v.Name,
				Position: v.Position,
				TeamId:   v.TeamId,
				PTS:      v.PTS,
				REB:      v.REB,
				AST:      v.AST,
				STL:      v.STL,
				BLK:      v.BLK,
				Actions:  []string{},
			})
		}

		sm.Teams[i].Name = team.Name
		sm.Teams[i].Players = convPlayers

	}

	return nil
}

func (sm *simulation) createRounds() {
	teamCount := len(sm.Teams)
	rounds := teamCount - 1
	matchCount := teamCount / 2
	sm.Rounds = []round{}

	for i := 0; i < rounds; i++ {
		matches := make([]match, matchCount)

		for j := 0; j < matchCount; j++ {
			home := (i + j) % (rounds)
			away := (rounds - j + i) % (rounds)

			if j == 0 {
				away = rounds
			}

			src := rand.NewSource(12345678)
			rnd := rand.New(src)

			tempAwayPlayers := make([]player, len(sm.Teams[away].Players))
			tempHomePlayers := make([]player, len(sm.Teams[home].Players))

			copy(tempAwayPlayers, sm.Teams[away].Players)
			copy(tempHomePlayers, sm.Teams[home].Players)

			matches[j] = match{
				Home: team{
					Name:    sm.Teams[home].Name,
					Players: tempHomePlayers,
					Points:  0,
				},
				Away: team{
					Name:    sm.Teams[away].Name,
					Players: tempAwayPlayers,
					Points:  0,
				},
				Random: rnd,
			}

			if i > rounds/2 {
				matches[j].Home = team{
					Name:    sm.Teams[away].Name,
					Players: tempAwayPlayers,
					Points:  0,
				}
				matches[j].Away = team{
					Name:    sm.Teams[home].Name,
					Players: tempHomePlayers,
					Points:  0,
				}
			}
		}

		sm.Rounds = append(sm.Rounds, round{
			Matches: matches,
		})
	}
}

func (sm *simulation) simulate(isRunWholeRounds bool) {
	sm.RuningRound = 0
	totalRound := sm.RuningRound + 1

	if isRunWholeRounds {
		totalRound = len(sm.Rounds)
	}

	finishTime := 48

	go func() {
		for i := sm.RuningRound; i < totalRound; i++ {
			fmt.Println("running round:")
			fmt.Println(sm.RuningRound)
			fmt.Println("total round:")
			fmt.Println(totalRound)

			if !sm.IsSimulating {
				break
			}

			sm.GameTime = gameTimes{
				oneMinute:        5,
				totalPassMinutes: 0,
				secondCounter:    0,
			}
			sm.Timer = time.NewTicker(1 * time.Second)
			for _ = range sm.Timer.C {
				if sm.GameTime.totalPassMinutes == finishTime {
					sm.Timer.Stop()
					fmt.Println("stopped")
					break
				}

				if sm.GameTime.secondCounter%5 == 0 {
					sm.GameTime.totalPassMinutes++
				}

				sm.runMatches()

				sm.GameTime.secondCounter++
				sm.RuningRound = i
			}
		}

		fmt.Println("simulation is ended")

		sm.IsSimulating = false
	}()
}

func (sm *simulation) runMatches() {
	currentRound := sm.Rounds[sm.RuningRound]

	for i := 0; i < len(currentRound.Matches); i++ {
		match := &currentRound.Matches[i]

		match.Random.Seed(time.Now().UnixNano())

		if match.AttackTimer == 24 {
			match.reverseSides()
			continue
		}

		isThereAnyAttackInTheseSecond := match.Random.NormFloat64()*10.0 + 5.0

		if isThereAnyAttackInTheseSecond < 0 {
			continue
		}

		randomPlayerIndex := match.Random.Intn(100)
		randomAwayPlayer := &match.Away.Players[randomPlayerIndex%5]
		randomHomePlayer := &match.Home.Players[randomPlayerIndex%5]

		canBlock := match.check(randomAwayPlayer.BLK, randomHomePlayer.BLK, -2, 0.1)

		if canBlock == 0 {
			fmt.Println(randomHomePlayer.Name + " blocked by " + randomAwayPlayer.Name)
			randomHomePlayer.Blocks++
		} else if canBlock == 1 {
			fmt.Println(randomAwayPlayer.Name + " blocked by " + randomHomePlayer.Name)
			randomAwayPlayer.Blocks++
		}

		if canBlock != -1 {
			match.reverseSides()
			continue
		}

		point := match.getPoint()

		if point != 0 {
			canPoint := match.check(randomAwayPlayer.PTS, randomHomePlayer.PTS, -0.01, 0.1)

			if canPoint == 0 {
				randomHomePlayerAssist := &match.Home.Players[match.Random.Intn(100)%5]
				canAssist := match.check(0, randomHomePlayerAssist.AST, -10, -1)

				if canAssist == 0 {
					fmt.Println(randomHomePlayerAssist.Name + " get assist")
					randomHomePlayerAssist.Assists++
				}

				fmt.Println(randomHomePlayer.Name + " get point")
				randomHomePlayer.Points += point
				randomHomePlayer.SuccessShoots++
				match.Home.Points += point
			} else if canPoint == 1 {
				randomAwayPlayerAssist := &match.Away.Players[match.Random.Intn(100)%5]
				canAssist := match.check(randomAwayPlayerAssist.AST, 0, -10, -1)

				if canAssist == 1 {
					fmt.Println(randomAwayPlayerAssist.Name + " get assist")
					randomAwayPlayerAssist.Assists++
				}

				fmt.Println(randomAwayPlayer.Name + " get point")
				randomAwayPlayer.Points += point
				randomAwayPlayer.SuccessShoots++
				match.Away.Points += point
			} else {
				if match.CurrentSide == 0 {
					randomHomePlayer.FailShoots++
				} else {
					randomAwayPlayer.FailShoots++
				}
			}

			match.reverseSides()
			continue
		}

		match.AttackTimer++
	}
}
