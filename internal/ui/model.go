package ui

import (
	"time"

	"github.com/BosBJJ/hjkl-hero/internal/game"
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
	Screen   Screen
	Menu     MenuModel
	Game     GameModel
	GameOver GameOverModel
}

func NewModel() Model {
	return Model{
		Menu:     MakeMenu(),
		Game:     NewGameModel(),
		GameOver: MakeGameOver(),
	}
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
				m.Screen = GameScreen
				m.Menu.Selected = -1
			case 3:
				return m, tea.Quit
			}
			return m, cmd
		case GameScreen:
			game, cmd := m.Game.Update(msg)
			m.Game = game
			if game.GameOver{
				m.Screen = GameOverScreen
			}
			return m, cmd
		case GameOverScreen:
			gameOver, cmd := m.GameOver.UpdateGameOver(msg)
			m.GameOver = gameOver
			switch gameOver.Selected {
			case 0:
				m.Screen = MainMenuScreen
				m.GameOver.Selected = -1
			case 1:
				m.Screen = MainMenuScreen
				m.GameOver.Selected = -1
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
	}
	return "No Screen Selected"
}
