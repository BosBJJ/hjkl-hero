package ui

import (
	"database/sql"
	"time"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	MainMenuScreen Screen = iota
	GameScreen
	GameOverScreen
	HighScoresScreen
	SettingsScreen
)

type tickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type Model struct {
	Screen     Screen
	Menu       MenuModel
	Game       GameModel
	GameOver   GameOverModel
	Settings   SettingsModel
	HighScores HighScoresModel
	DB         *sql.DB
}

func NewModel(db *sql.DB) Model {
	return Model{
		Menu:       MakeMenu(),
		Game:       MakeDefaultGameModel(),
		GameOver:   MakeGameOver(),
		Settings:   MakeSettingsModel(),
		HighScores: MakeHighScores(),
		DB:         db,
	}
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		menu, _ := m.Menu.UpdateMenu(msg)
		m.Menu = menu

		game, _ := m.Game.Update(msg)
		m.Game = game

		gameOver, _ := m.GameOver.UpdateGameOver(msg)
		m.GameOver = gameOver

		optMenu, _ := m.Settings.UpdateSettings(msg)
		m.Settings = optMenu

		hs, _ := m.HighScores.UpdateHighScores(msg)
		m.HighScores = hs

		return m, nil
	case tickMsg:
		if m.Game.gameState.MapInfo.MapType == game.RoomMap {
			m.Game.gameState.SpawnEnemy()
		}
		return m, doTick()
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		}
		switch m.Screen {
		case MainMenuScreen:
			menu, cmd := m.Menu.UpdateMenu(msg)
			m.Menu = menu
			switch menu.Selected {
			case 0:
				m.Game = MakeDefaultGameModel()
				m.Screen = GameScreen
				m.Menu.Selected = -1
			case 1:
				hs, _ := storage.ShowScores(m.DB)
				m.HighScores.Scores = hs
				m.Screen = HighScoresScreen
				m.Menu.Selected = -1
			case 2:
				m.Screen = SettingsScreen
				m.Menu.Selected = -1
			case 3:
				return m, tea.Quit
			}
			return m, cmd
		case GameScreen:
			game, cmd := m.Game.Update(msg)
			m.Game = game
			if game.GameOver {
				m.GameOver.Stats = GetRunStats(m.Game)
				m.Screen = GameOverScreen
			}
			return m, cmd
		case GameOverScreen:
			gameOver, cmd := m.GameOver.UpdateGameOver(msg)
			m.GameOver = gameOver
			switch gameOver.Selected {
			case 0:
				storage.SaveRun(m.DB, storage.Run{
					PlayerName: gameOver.PlayerName,
					Kills:      gameOver.Stats.Kills,
					TotalXp:    gameOver.Stats.TotalXp,
					TotalMoves: gameOver.Stats.TotalMoves,
					MapLevel:   gameOver.Stats.MapLevel,
					GameMode:   storage.TutorialMode,
				})
				hs, _ := storage.ShowScores(m.DB)
				m.HighScores.Scores = hs
				m.Screen = HighScoresScreen
				m.GameOver.Selected = -1
			case 1:
				m.Screen = MainMenuScreen
				m.GameOver.Selected = -1
			}
			return m, cmd
		case SettingsScreen:
			settings, cmd := m.Settings.UpdateSettings(msg)
			m.Settings = settings
			switch settings.Selected {
			case 2:
				m.Screen = MainMenuScreen
				m.Settings.Selected = -1
			}
			return m, cmd
		case HighScoresScreen:
			hs, cmd := m.HighScores.UpdateHighScores(msg)
			m.HighScores = hs
			switch hs.Selected {
			case 0:
				m.Screen = MainMenuScreen
				m.HighScores.Selected = -1
			}
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.Screen {
	case MainMenuScreen:
		return m.Menu.ViewMenu()
	case GameScreen:
		return m.Game.ViewGame()
	case GameOverScreen:
		return m.GameOver.ViewGameOver()
	case SettingsScreen:
		return m.Settings.ViewSettings()
	case HighScoresScreen:
		return m.HighScores.ViewHighScores()
	}
	return "No Screen Selected"
}
