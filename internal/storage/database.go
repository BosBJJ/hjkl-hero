package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Run struct {
	Id         int
	PlayerName string
	Kills      int
	TotalXp    int
	TotalMoves int
	MapLevel   int
	GameMode   GameMode
	FinishedAt string
}

type Settings struct {
	Theme    Theme
	GameMode GameMode
}

func MakeDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database at %s: %v", path, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error creating DB: %v", err)
	}
	return db, nil
}

func CreateHSSchema(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS scores(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	kills INTEGER,
	total_xp INTEGER,
	total_moves INTEGER,
	map_level INTEGER,
	game_mode TEXT,
	finished_at TEXT);
	`
	_, err := db.Exec(query)
	return err
}

func CreateSettingSchema(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS settings(
	id INTEGER PRIMARY KEY CHECK (id = 1),
	theme TEXT NOT NULL,
	game_mode TEXT NOT NULL);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = `INSERT OR IGNORE INTO settings (id, theme, game_mode)
	VALUES (1, 'default', 'tutorial');
	`

	_, err = db.Exec(query)

	return err
}

func GetSettings(db *sql.DB) (Settings, error) {
	query := `SELECT theme, game_mode FROM settings`

	var settings Settings
	err := db.QueryRow(query).Scan(
		&settings.Theme, &settings.GameMode)

	if err != nil {
		return Settings{}, fmt.Errorf("error: %v", err)
	}

	return settings, nil
}

func UpdateGameMode(db *sql.DB, mode GameMode) error {
	_, err := db.Exec(`
        UPDATE settings
        SET game_mode = ?
        WHERE id = 1
    `, mode)

	return err
}

func UpdateTheme(db *sql.DB, theme Theme) error {
	_, err := db.Exec(`
        UPDATE settings
        SET theme = ?
        WHERE id = 1
    `, theme)

	return err
}

func SaveRun(db *sql.DB, run Run) error {
	run.FinishedAt = time.Now().UTC().Format(time.RFC3339)

	query := `INSERT INTO scores (name, kills, total_xp, total_moves, map_level, game_mode, finished_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query, run.PlayerName, run.Kills, run.TotalXp, run.TotalMoves, run.MapLevel, run.GameMode, run.FinishedAt)

	return err
}

func ShowScores(db *sql.DB) ([]Run, error) {
	query := `SELECT * FROM scores`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error :%v", err)
	}
	defer rows.Close()

	var runs []Run

	for rows.Next() {
		var run Run
		err = rows.Scan(&run.Id, &run.PlayerName, &run.Kills, &run.TotalXp, &run.TotalMoves, &run.MapLevel, &run.GameMode, &run.FinishedAt)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		runs = append(runs, run)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return runs, nil
}
