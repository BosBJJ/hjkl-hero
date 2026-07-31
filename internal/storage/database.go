package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Run struct {
	Id          int
	PlayerName  string
	Kills       int
	TotalXp     int
	TotalMoves  int
	MapLevel    int
	GameMode    GameMode
	FinishedAt  string
	DamageTaken int
}

type Settings struct {
	ThemeID  ThemeID
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
	finished_at TEXT,
	damage_taken INTEGER);
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

func CreateCustomTheme(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS custom_theme(
	id INTEGER PRIMARY KEY CHECK (id = 1),
	wall_color TEXT NOT NULL,
	floor_color TEXT NOT NULL,
	player_color TEXT NOT NULL,
	wall_icon TEXT NOT NULL,
	floor_icon TEXT NOT NULL,
	player_icon TEXT NOT NULL);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = `INSERT OR IGNORE INTO custom_theme (id, wall_color, floor_color, player_color, wall_icon, floor_icon, player_icon)
	VALUES (?,?,?,?,?,?,?);
	`

	_, err = db.Exec(query, 1, "magenta", "green", "red", "#", ".", "#")

	return err
}

func GetSettings(db *sql.DB) (Settings, error) {
	query := `SELECT theme, game_mode FROM settings`

	var settings Settings
	err := db.QueryRow(query).Scan(
		&settings.ThemeID, &settings.GameMode)

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

func UpdateTheme(db *sql.DB, theme ThemeID) error {
	_, err := db.Exec(`
        UPDATE settings
        SET theme = ?
        WHERE id = 1
    `, theme)

	return err
}

func SaveCustomTheme(db *sql.DB, theme Theme) error {
	_, err := db.Exec(`
        UPDATE custom_theme
		SET
			wall_color = ?,
			floor_color = ?,
			player_color = ?,
			wall_icon = ?,
			floor_icon = ?,
			player_icon = ?
        WHERE id = 1
    `, theme.WallColor, theme.FloorColor, theme.PlayerColor, theme.WallIcon, theme.FloorIcon, theme.PlayerIcon)
	return err
}

func LoadCustomTheme(db *sql.DB) (Theme, error) {
	query := `SELECT wall_color, floor_color, player_color, wall_icon, floor_icon, player_icon FROM custom_theme
	WHERE id = 1`

	var theme Theme
	err := db.QueryRow(query).Scan(
		&theme.WallColor, &theme.FloorColor, &theme.PlayerColor, &theme.WallIcon, &theme.FloorIcon, &theme.PlayerIcon)

	if err != nil {
		return Theme{}, fmt.Errorf("error: %v", err)
	}

	return theme, nil
}

func SaveRun(db *sql.DB, run Run) error {
	run.FinishedAt = time.Now().UTC().Format(time.RFC3339)

	query := `INSERT INTO scores (name, kills, total_xp, total_moves, map_level, game_mode, finished_at, damage_taken)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query, run.PlayerName, run.Kills, run.TotalXp, run.TotalMoves, run.MapLevel, run.GameMode, run.FinishedAt, run.DamageTaken)

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
		err = rows.Scan(&run.Id, &run.PlayerName, &run.Kills, &run.TotalXp, &run.TotalMoves, &run.MapLevel, &run.GameMode, &run.FinishedAt, &run.DamageTaken)
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
