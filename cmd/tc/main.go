package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

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
	app.SetWorkspace(resolveWorkspace())
	app.SetProviderService(providerSvc, modelRepo, sessionRepo, messageRepo)

	program := tea.NewProgram(app)
	app.SetProgram(program)

	if _, err := program.Run(); err != nil {
		logger.Error("tui error", "error", err)
		os.Exit(1)
	}
}

func resolveWorkspace() string {
	args := os.Args[1:]
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		path := args[0]
		if strings.HasPrefix(path, "~/") {
			home, _ := os.UserHomeDir()
			path = filepath.Join(home, path[2:])
		}
		abs, err := filepath.Abs(path)
		if err == nil {
			return abs
		}
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "~"
	}
	return cwd
}
