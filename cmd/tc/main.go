package main

import (
	"log/slog"
	"os"

	tea "charm.land/bubbletea/v2"

	"termcode/internal/adapters/tui"
	"termcode/internal/application/provider"
	sqliterepo "termcode/internal/infrastructure/database/sqlite"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))

	db, err := sqliterepo.Open()
	if err != nil {
		logger.Error("database open", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := sqliterepo.RunMigrations(db); err != nil {
		logger.Error("database migrate", "error", err)
		os.Exit(1)
	}

	providerRepo := sqliterepo.NewProviderRepo(db)
	modelRepo := sqliterepo.NewModelRepo(db)
	sessionRepo := sqliterepo.NewSessionRepo(db)
	messageRepo := sqliterepo.NewMessageRepo(db)

	providerSvc := provider.NewService(providerRepo, logger)

	app := tui.NewApp()
	app.SetProviderService(providerSvc, modelRepo, sessionRepo, messageRepo)

	program := tea.NewProgram(app)
	app.SetProgram(program)

	if _, err := program.Run(); err != nil {
		logger.Error("tui error", "error", err)
		os.Exit(1)
	}
}
