package repository

import (
	"nba_stats/model"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	DB *sqlx.DB
}

func NewRepository() *repository {
	db, err := sqlx.Connect("mysql", "root:nba_stats@tcp(db:3306)/nba")

	if err != nil {
		panic(err)
	}

	return &repository{
		DB: db,
	}
}

func (r *repository) GetTeams() ([]model.Team, error) {
	teams := []model.Team{}
	err := r.DB.Select(&teams, model.SELECT_TEAM_SQL)

	return teams, err
}

func (r *repository) GetPlayers(teamId int) ([]model.Player, error) {
	players := []model.Player{}
	err := r.DB.Select(&players, model.SELECT_PLAYER_SQL, teamId)

	return players, err
}

func (r *repository) SaveTeam(team model.Team) {
	r.DB.Exec(model.INSERT_TEAM_SQL, team.Name)
}

func (r *repository) SavePlayer(player model.Player) {
	r.DB.Exec(model.INSERT_PLAYER_SQL, player.Name, player.Position, player.TeamId, player.PTS, player.REB, player.AST, player.STL, player.BLK)
}

func (r *repository) RemoveTeams() {
	r.DB.Exec(model.DELETE_ALL_TEAM)
}

func (r *repository) RemovePlayer() {
	r.DB.Exec(model.DELETE_ALL_PLAYERS)
}

func (r *repository) CreatePlayersTable() {
	r.DB.Exec(model.CREATE_PLAYERS_TABLE)
}

func (r *repository) CreateTeamsTable() {
	r.DB.Exec(model.CREATE_TEAMS_TABLE)
}
