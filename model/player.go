package model

type Player struct {
	Id       int     `db:"id"`
	Name     string  `db:"name"`
	Position string  `db:"position"`
	TeamId   int     `db:"team_id"`
	PTS      float64 `db:"pts"`
	REB      float64 `db:"reb"`
	AST      float64 `db:"ast"`
	STL      float64 `db:"stl"`
	BLK      float64 `db:"blk"`
}

const CREATE_PLAYERS_TABLE = `
CREATE TABLE IF NOT EXISTS players (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(150) NULL DEFAULT NULL,
	position VARCHAR(10) NULL DEFAULT NULL,
	team_id INT NULL DEFAULT NULL,
	pts FLOAT NULL DEFAULT NULL,
	reb FLOAT NULL DEFAULT NULL,
	ast FLOAT NULL DEFAULT NULL,
	stl FLOAT NULL DEFAULT NULL,
	blk FLOAT NULL DEFAULT NULL,
	PRIMARY KEY (id)
)
COLLATE='utf8_general_ci';
`
const INSERT_PLAYER_SQL = "INSERT INTO `players` (`name`, `position`, `team_id`, `pts`, `reb`, `ast`, `stl`, `blk`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
const SELECT_PLAYER_SQL = "SELECT * FROM `players` WHERE `team_id`=?;"
const DELETE_ALL_PLAYERS = "DELETE FROM `players`"
