package model

type Team struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

const CREATE_TEAMS_TABLE = `
CREATE TABLE IF NOT EXISTS teams (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(100) NULL DEFAULT NULL,
	PRIMARY KEY (id)
)
COLLATE='utf8_general_ci';
`
const INSERT_TEAM_SQL = "INSERT INTO `teams` (`name`) VALUES (?);"
const SELECT_TEAM_SQL = "SELECT * FROM `teams`;"
const DELETE_ALL_TEAM = "DELETE FROM `teams`;"
