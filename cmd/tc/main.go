package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/spf13/cobra"

	"termcode/internal/adapters/tui"
	"termcode/internal/application/provider"
	sqliterepo "termcode/internal/infrastructure/database/sqlite"
)

var (
	Version   = "0.1.0"
	Commit    = "dev"
	workspace string
	logger    *slog.Logger
)

func main() {
	logger = newLogger()

	rootCmd := &cobra.Command{
		Use:   "tc",
		Short: "TermCode - AI Coding CLI",
		Long:  "TermCode is a mobile-first AI coding CLI for Termux.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			w, _ := cmd.Flags().GetString("workspace")
			workspace = w
			if workspace == "" {
				cwd, _ := os.Getwd()
				workspace = cwd
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTUI()
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(chatCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(providerCmd)

	rootCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "", "workspace directory")
	rootCmd.Flags().BoolP("version", "v", false, "version for tc")

	_ = rootCmd.Execute()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TermCode v%s (commit: %s)\n", Version, Commit)
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat [workspace]",
	Short: "Start interactive chat",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			workspace = args[0]
		}
		return runTUI()
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage AI providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	configCmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Workspace: %s\n", workspace)
			return nil
		},
	})
}

func runTUI() error {
	db, err := openDB()
	if err != nil {
		logger.Error("database open", "error", err)
		return err
	}
	defer db.Close()

	if err := runMigrations(db); err != nil {
		logger.Error("database migrate", "error", err)
		return err
	}

	providerRepo := openProviderRepo(db)
	modelRepo := sqliterepo.NewModelRepo(db)
	sessionRepo := sqliterepo.NewSessionRepo(db)
	messageRepo := sqliterepo.NewMessageRepo(db)
	settingsRepo := sqliterepo.NewSettingsRepo(db)
	branchRepo := sqliterepo.NewBranchRepo(db)

	providerSvc := provider.NewService(providerRepo, logger)

	app := tui.NewApp()
	app.SetWorkspace(workspace)
	app.SetProviderService(providerSvc, modelRepo, sessionRepo, messageRepo, settingsRepo, branchRepo)

	program := tea.NewProgram(app, tea.WithFPS(30))
	app.SetProgram(program)

	if _, err := program.Run(); err != nil {
		logger.Error("tui error", "error", err)
		return err
	}
	return nil
}

func newLogger() *slog.Logger {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	logDir := filepath.Join(homeDir, ".config", "termcode")
	os.MkdirAll(logDir, 0755)
	logPath := filepath.Join(logDir, "termcode.log")
	if fi, _ := os.Stat(logPath); fi != nil && fi.Size() > 10*1024*1024 {
		os.Remove(logPath)
	}
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
