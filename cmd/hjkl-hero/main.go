package main

import (
	"log"

	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	db, err := storage.MakeDB("hjkl-hero.db")
	if err != nil {
		log.Fatalf("unable to create database: %v", err)
	}
	defer db.Close()

	err = storage.CreateHSSchema(db)
	if err != nil {
		log.Fatalf("unable to create database: %v", err)
	}

	p := tea.NewProgram(ui.NewModel(db), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
