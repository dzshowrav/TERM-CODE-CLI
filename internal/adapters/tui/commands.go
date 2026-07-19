package tui

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	"termcode/internal/adapters/tui/screens"
	"termcode/internal/application/model"
	"termcode/internal/application/provider"
	domainmodel "termcode/internal/domain/model"
	"termcode/internal/domain/session"
	"termcode/internal/infrastructure/collab"
	"termcode/internal/infrastructure/eventbus"
	git "termcode/internal/infrastructure/git"
)

type cmdHandler func(ctx context.Context, args []string) string

type commandRegistry struct {
	handlers map[string]cmdHandler
	app      *AppModel
}

func newCommandRegistry(providerSvc *provider.Service, modelSvc *model.Service, app *AppModel) *commandRegistry {
	r := &commandRegistry{
		handlers: make(map[string]cmdHandler),
		app:      app,
	}

	r.register("help", r.cmdHelp)
	r.register("h", r.cmdHelp)
	r.register("provider", func(ctx context.Context, args []string) string {
		return r.cmdProvider(ctx, args, providerSvc, modelSvc)
	})
	r.register("providers", func(ctx context.Context, args []string) string {
		return r.cmdProvider(ctx, nil, providerSvc, modelSvc)
	})
	r.register("models", func(ctx context.Context, args []string) string {
		return r.cmdModels(ctx, args, modelSvc)
	})
	r.register("model", func(ctx context.Context, args []string) string {
		return r.cmdModels(ctx, args, modelSvc)
	})
	r.register("addmodel", func(ctx context.Context, args []string) string {
		return r.cmdAddModel(ctx, args, providerSvc, modelSvc)
	})
	r.register("agents", r.cmdAgent)
	r.register("agent", r.cmdAgent)
	r.register("workspace", r.cmdWorkspace)
	r.register("sessions", r.cmdSessions)
	r.register("session", r.cmdSessions)
	r.register("clear", r.cmdClear)
	r.register("tools", r.cmdTools)
	r.register("tool", r.cmdTools)
	r.register("git", r.cmdGit)
	r.register("settings", r.cmdSettings)
	r.register("exit", r.cmdExit)
	r.register("quit", r.cmdExit)
	r.register("home", r.cmdHome)
	r.register("about", r.cmdAbout)
	r.register("network", r.cmdNetwork)
	r.register("status", r.cmdNetwork)
	r.register("batch", r.cmdBatch)
	r.register("stop", r.cmdStop)
	r.register("continue", r.cmdContinue)
	r.register("edit", r.cmdEdit)
	r.register("delete", r.cmdDelete)
	r.register("rename", r.cmdRename)
	r.register("branch", r.cmdBranch)
	r.register("undo", r.cmdUndo)
	r.register("redo", r.cmdRedo)
	r.register("export", r.cmdExport)
	r.register("import", r.cmdImport)
	r.register("pin", r.cmdPin)
	r.register("search", r.cmdSearch)
	r.register("collab", r.cmdCollab)
	r.register("retry", r.cmdRetry)
	r.register("cancel", r.cmdCancel)
	r.register("theme", r.cmdTheme)
	r.register("plugin", r.cmdPlugin)

	return r
}

func (r *commandRegistry) register(name string, handler cmdHandler) {
	r.handlers[name] = handler
}

func (r *commandRegistry) execute(ctx context.Context, input string) string {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "/") {
		return ""
	}

	parts := splitArgs(input[1:])
	if len(parts) == 0 {
		return "Type /help for available commands."
	}

	cmd := parts[0]
	args := parts[1:]

	handler, ok := r.handlers[cmd]
	if !ok {
		return fmt.Sprintf("Unknown command: /%s. Type /help for available commands.", cmd)
	}

	return handler(ctx, args)
}

func (r *commandRegistry) cmdHelp(ctx context.Context, args []string) string {
	r.app.ShowDialog(screens.NewHelpScreen())
	return "__dialog__"
}

func (r *commandRegistry) showMessageDialog(title, content string) string {
	r.app.ShowDialog(screens.NewMessageScreen(title, content))
	return "__dialog__"
}

func (r *commandRegistry) cmdProvider(ctx context.Context, args []string, providerSvc *provider.Service, modelSvc *model.Service) string {
	if len(args) > 0 && args[0] == "add" {
		if len(args) >= 3 {
			name := args[1]
			baseURL := args[2]
			apiKey := ""
			if len(args) > 3 {
				apiKey = args[3]
			}
			p, err := providerSvc.Create(ctx, name, baseURL, apiKey, "")
			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Provider '%s' added (ID: %s)", p.Name, p.ID)
		}
		s := screens.NewProviderAddScreen(
			func(name, baseURL, apiKey, desc string) string {
				p, err := providerSvc.Create(ctx, name, baseURL, apiKey, desc)
				if err != nil {
					return fmt.Sprintf("Error: %v", err)
				}
				return fmt.Sprintf("Provider '%s' added (ID: %s)", p.Name, p.ID)
			},
			func(name, baseURL, apiKey, desc string) string {
				p, err := providerSvc.Create(ctx, name, baseURL, apiKey, desc)
				if err != nil {
					return fmt.Sprintf("Error: %v", err)
				}
				result, err := providerSvc.TestConnection(ctx, p.ID)
				if err != nil {
					return fmt.Sprintf("Test error: %v", err)
				}
				if result.Success {
					return fmt.Sprintf("OK (%dms, %d models)", result.Latency, result.Models)
				}
				return fmt.Sprintf("Failed: %s", result.Message)
			},
		)
		r.app.ShowDialog(s)
		return "__dialog__"
	}

	if len(args) > 0 && args[0] == "sync" {
		p, err := providerSvc.GetDefault(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		apiKey, _ := providerSvc.DecryptAPIKey(p.APIKey)
		n, err := modelSvc.SyncFromProvider(ctx, p.ID, p.BaseURL, apiKey)
		if err != nil {
			return fmt.Sprintf("Error syncing models: %v", err)
		}
		return fmt.Sprintf("Synced %d models from '%s'.", n, p.Name)
	}

	providers, err := providerSvc.List(ctx)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	items := make([]screens.ProviderItem, 0, len(providers))
	defID := ""
	for _, p := range providers {
		if p.IsDefault {
			defID = p.ID
		}
	}
	for _, p := range providers {
		status := "disconnected"
		items = append(items, screens.ProviderItem{
			Name:     p.Name,
			URL:      p.BaseURL,
			Status:   status,
			Latency:  0,
			IsActive: p.ID == defID,
		})
	}
	s := screens.NewProviderListScreen()
	s.SetProviders(items)
	providersCopy := providers
	s.OnRefresh(func() {
		refreshed, err := providerSvc.List(ctx)
		if err != nil {
			return
		}
		defID := ""
		for _, p := range refreshed {
			if p.IsDefault {
				defID = p.ID
			}
		}
		newItems := make([]screens.ProviderItem, 0, len(refreshed))
		for _, p := range refreshed {
			status := "disconnected"
			newItems = append(newItems, screens.ProviderItem{
				Name:     p.Name,
				URL:      p.BaseURL,
				Status:   status,
				Latency:  0,
				IsActive: p.ID == defID,
			})
		}
		providersCopy = refreshed
		s.SetProviders(newItems)
	})
	// Note: providersCopy becomes stale after delete, but dialog closes after onDelete so it's OK
	s.OnDelete(func(name string) string {
		for _, p := range providersCopy {
			if strings.EqualFold(p.Name, name) {
				if r.app.modelRepo != nil {
					if err := r.app.modelRepo.DeleteByProvider(ctx, p.ID); err != nil {
						return fmt.Sprintf("Error deleting models: %v", err)
					}
				}
				if err := providerSvc.Delete(ctx, p.ID); err != nil {
					return fmt.Sprintf("Error: %v", err)
				}
				return fmt.Sprintf("Provider '%s' and its models deleted.", name)
			}
		}
		return fmt.Sprintf("Provider '%s' not found.", name)
	})
	s.OnEdit(func(name string) string {
		for _, p := range providersCopy {
			if strings.EqualFold(p.Name, name) {
				apiKey, _ := providerSvc.DecryptAPIKey(p.APIKey)
				editScreen := screens.NewProviderEditScreen(
					p.ID, p.Name, p.BaseURL, apiKey, p.Description,
					func(id, name, baseURL, apiKey, desc string) string {
						_, err := providerSvc.Update(ctx, id, name, baseURL, apiKey, desc)
						if err != nil {
							return fmt.Sprintf("Error: %v", err)
						}
						return fmt.Sprintf("Provider '%s' updated.", name)
					},
					func(id, name, baseURL, apiKey, desc string) string {
						result, err := providerSvc.TestConnection(ctx, id)
						if err != nil {
							return fmt.Sprintf("Test error: %v", err)
						}
						if result.Success {
							return fmt.Sprintf("OK (%dms, %d models)", result.Latency, result.Models)
						}
						return fmt.Sprintf("Failed: %s", result.Message)
					},
				)
				r.app.ShowDialog(editScreen)
				return "__dialog__"
			}
		}
		return fmt.Sprintf("Provider '%s' not found.", name)
	})
	s.OnSelect(func(name string) string {
		for _, p := range providersCopy {
			if strings.EqualFold(p.Name, name) {
				if err := providerSvc.SetDefault(ctx, p.ID); err != nil {
					return fmt.Sprintf("Error: %v", err)
				}
				r.app.providerName = p.Name
				r.app.homeScreen.UpdateConfig(screens.HomeScreenConfig{ProviderName: p.Name, ProviderURL: p.BaseURL})
				r.app.statusBar.SetModel(p.Name + "/" + r.app.modelName)
				if r.app.settingsRepo != nil {
					r.app.settingsRepo.Set(ctx, "provider_id", p.ID)
				}
				return fmt.Sprintf("Provider '%s' is now active.", name)
			}
		}
		return fmt.Sprintf("Provider '%s' not found.", name)
	})
	r.app.ShowDialog(s)
	return "__dialog__"
}

func (r *commandRegistry) cmdModels(ctx context.Context, args []string, modelSvc *model.Service) string {
	return r.showModelSelectDialog(ctx, modelSvc)
}

func (r *commandRegistry) showModelSelectDialog(ctx context.Context, modelSvc *model.Service) string {
	pmodel, err := modelSvc.List(ctx)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	models := make([]domainmodel.Model, len(pmodel))
	for i, m := range pmodel {
		models[i] = *m
	}
	s := screens.NewModelSelectScreen(func(modelID string) string {
		r.app.SetSelectedModel(modelID)
		return fmt.Sprintf("Model '%s' selected.", modelID)
	})
	s.SetModels(models)
	r.app.ShowDialog(s)
	return "__dialog__"
}

func (r *commandRegistry) cmdAddModel(ctx context.Context, args []string, providerSvc *provider.Service, modelSvc *model.Service) string {
	s := screens.NewModelAddScreen(func(id, display, provider string, ctxSize, maxOut int) string {
		m := &domainmodel.Model{
			ModelID:     id,
			DisplayName: display,
			MaxContext:  ctxSize,
			MaxOutput:   maxOut,
		}
		if err := modelSvc.Create(ctx, m); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return fmt.Sprintf("Model '%s' added.", display)
	})
	r.app.ShowDialog(s)
	return "__dialog__"
}

func (r *commandRegistry) cmdAgent(ctx context.Context, args []string) string {
	s := screens.NewAgentSelectScreen(func(name string) string {
		r.app.agentName = name
		r.app.homeScreen.UpdateConfig(screens.HomeScreenConfig{AgentName: name})
		r.app.statusBar.SetAgent(name)
		if r.app.settingsRepo != nil {
			r.app.settingsRepo.Set(ctx, "agent_name", name)
		}
		return ""
	})
	r.app.ShowDialog(s)
	return "__dialog__"
}

func (r *commandRegistry) cmdWorkspace(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return r.showMessageDialog("Workspace", fmt.Sprintf("Current workspace:\n  %s", r.app.workspace))
	}
	path := strings.Join(args, " ")
	r.app.workspace = path
	r.app.homeScreen.UpdateConfig(screens.HomeScreenConfig{WorkspacePath: path})
	return fmt.Sprintf("Workspace set to '%s'.", path)
}

func (r *commandRegistry) cmdSessions(ctx context.Context, args []string) string {
	if r.app.sessionRepo == nil {
		return "Session storage not available."
	}

	if len(args) > 0 && args[0] == "new" {
		r.app.history = nil
		r.app.streamBuf = ""
		r.app.currentSess = nil
		r.app.chatScreen = screens.NewChatScreen()
		r.app.screen = screenHome
		return "__home__"
	}

	sessions, err := r.app.sessionRepo.ListActive(ctx)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	items := make([]screens.SessionListItem, 0, len(sessions))
	for _, s := range sessions {
		items = append(items, screens.SessionListItem{
			ID:       s.ID,
			Name:     s.Name,
			MsgCount: s.MessageCnt,
			IsActive: r.app.currentSess != nil && s.ID == r.app.currentSess.ID,
		})
	}
	ss := screens.NewSessionScreen()
	ss.SetSessions(items)
	ss.OnDelete(func(id string) string {
		if r.app.messageRepo != nil {
			r.app.messageRepo.DeleteBySession(ctx, id)
		}
		if err := r.app.sessionRepo.Delete(ctx, id); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return fmt.Sprintf("Session '%s' deleted.", id)
	})
	r.app.ShowDialog(ss)
	return "__dialog__"
}

func (r *commandRegistry) cmdTools(ctx context.Context, args []string) string {
	if r.app.chatSvc == nil {
		return "Service not initialized."
	}
	tools := r.app.chatSvc.Tools()
	s := screens.NewToolListScreen(tools)
	r.app.ShowDialog(s)
	return "__dialog__"
}

func (r *commandRegistry) cmdGit(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return r.showMessageDialog("Git", "Usage: /git status  |  /git log  |  /git diff  |  /git add  |  /git commit  |  /git branches")
	}

	svc := git.NewService()
	if !svc.IsRepo(r.app.workspace) {
		return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
	}
	repo, err := svc.Open(r.app.workspace)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	switch args[0] {
	case "status":
		status, err := svc.Status(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if status.Clean {
			return r.showMessageDialog("Git Status", fmt.Sprintf("On branch %s. Clean working tree.", status.Branch))
		}
		var parts []string
		parts = append(parts, fmt.Sprintf("On branch %s (%s)", status.Branch, status.Hash))
		if len(status.Staged) > 0 {
			parts = append(parts, "")
			parts = append(parts, "Staged:")
			for _, f := range status.Staged {
				parts = append(parts, fmt.Sprintf("  %s", f))
			}
		}
		if len(status.Modified) > 0 {
			parts = append(parts, "")
			parts = append(parts, "Modified:")
			for _, f := range status.Modified {
				parts = append(parts, fmt.Sprintf("  %s", f))
			}
		}
		if len(status.Added) > 0 {
			parts = append(parts, "")
			parts = append(parts, "Added:")
			for _, f := range status.Added {
				parts = append(parts, fmt.Sprintf("  %s", f))
			}
		}
		if len(status.Deleted) > 0 {
			parts = append(parts, "")
			parts = append(parts, "Deleted:")
			for _, f := range status.Deleted {
				parts = append(parts, fmt.Sprintf("  %s", f))
			}
		}
		return r.showMessageDialog("Git Status", strings.Join(parts, "\n"))

	case "log":
		count := 10
		if len(args) > 1 {
			fmt.Sscanf(args[1], "%d", &count)
		}
		entries, err := svc.Log(repo, count)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		var lines []string
		for _, e := range entries {
			msg := strings.SplitN(e.Message, "\n", 2)[0]
			lines = append(lines, fmt.Sprintf("%s  %s  %s", e.When.Format("Jan 02 15:04"), e.Hash, msg))
		}
		return r.showMessageDialog("Git Log", strings.Join(lines, "\n"))

	case "diff":
		d, err := svc.Diff(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if len(d.Files) == 0 {
			return r.showMessageDialog("Git Diff", "No changes.")
		}
		files := make([]screens.DiffFile, 0, len(d.Files))
		for _, f := range d.Files {
			files = append(files, screens.DiffFile{
				Path:    f.Name,
				Status:  "modified",
				Added:   f.Added,
				Removed: f.Removed,
				Content: f.Content,
			})
		}
		r.app.ShowDialog(screens.NewDiffScreen(files))
		return "__dialog__"

	case "add":
		status, err := svc.Status(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}

		var unstaged []screens.GitFileItem
		for _, f := range status.Modified {
			unstaged = append(unstaged, screens.GitFileItem{Name: f, Status: "M"})
		}
		for _, f := range status.Added {
			unstaged = append(unstaged, screens.GitFileItem{Name: f, Status: "A"})
		}
		for _, f := range status.Deleted {
			unstaged = append(unstaged, screens.GitFileItem{Name: f, Status: "D"})
		}

		if len(unstaged) == 0 {
			if len(status.Staged) > 0 {
				return "All changes already staged."
			}
			return r.showMessageDialog("Git Add", "No unstaged files.")
		}

		s := screens.NewGitAddScreen(unstaged, func(files []string) string {
			if err := svc.Add(repo, files); err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Staged %d file(s): %s", len(files), strings.Join(files, ", "))
		})
		r.app.ShowDialog(s)
		return "__dialog__"

	case "commit":
		s := screens.NewGitCommitScreen(func(msg string) string {
			hash, err := svc.Commit(repo, msg)
			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			firstLine := strings.SplitN(msg, "\n", 2)[0]
			return fmt.Sprintf("Committed as %s: %s", hash, firstLine)
		})
		r.app.ShowDialog(s)
		return "__dialog__"

	case "branches":
		branches, err := svc.Branches(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		currentBranch, err := svc.GetBranch(repo)
		if err != nil {
			currentBranch = ""
		}
		var lines []string
		for _, b := range branches {
			if b == currentBranch {
				lines = append(lines, fmt.Sprintf("* %s (current)", b))
			} else {
				lines = append(lines, fmt.Sprintf("  %s", b))
			}
		}
		return r.showMessageDialog("Git Branches", strings.Join(lines, "\n"))

	case "checkout":
		if len(args) < 2 {
			return "Usage: /git checkout <branch>"
		}
		branch := args[1]
		if err := svc.Checkout(repo, branch, false); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		r.app.detectGitBranch()
		r.app.statusBar.SetBranch(r.app.gitBranch)
		return fmt.Sprintf("Switched to branch '%s'.", branch)

	default:
		return "Usage: /git status|log|diff|add|commit|branches|checkout"
	}
}

func (r *commandRegistry) cmdSettings(ctx context.Context, args []string) string {
	d := screens.NewSettingsScreen()

	var entries []screens.SettingEntry

	// Add current active config
	entries = append(entries, screens.SettingEntry{Key: "Active Provider", Value: r.app.providerName, Group: "Active Configuration"})
	entries = append(entries, screens.SettingEntry{Key: "Active Model", Value: r.app.modelName, Group: "Active Configuration"})
	entries = append(entries, screens.SettingEntry{Key: "Active Agent", Value: r.app.agentName, Group: "Active Configuration"})
	entries = append(entries, screens.SettingEntry{Key: "Workspace Path", Value: r.app.workspace, Group: "Active Configuration"})
	if r.app.gitBranch != "" {
		entries = append(entries, screens.SettingEntry{Key: "Git Branch", Value: r.app.gitBranch, Group: "Active Configuration"})
	}

	// Add database settings if available
	if r.app.settingsRepo != nil {
		dbSettings, err := r.app.settingsRepo.List(ctx)
		if err != nil {
			r.app.toastManager.Error("Cannot load settings: " + err.Error())
		} else {
			for k, v := range dbSettings {
				entries = append(entries, screens.SettingEntry{Key: k, Value: v, Group: "Database Settings"})
			}
		}
	}

	d.SetSettings(entries)
	r.app.ShowDialog(d)
	return "__dialog__"
}

func (r *commandRegistry) cmdClear(ctx context.Context, args []string) string {
	r.app.history = nil
	r.app.currentSess = nil
	if r.app.chatScreen != nil {
		r.app.chatScreen = screens.NewChatScreen()
	}
	return "__clear__"
}

func (r *commandRegistry) cmdExit(ctx context.Context, args []string) string {
	return "__exit__"
}

func (r *commandRegistry) cmdAbout(ctx context.Context, args []string) string {
	r.app.ShowDialog(screens.NewAboutScreen())
	return "__dialog__"
}

func (r *commandRegistry) cmdNetwork(ctx context.Context, args []string) string {
	state := screens.NetworkState{
		Status:      screens.NetworkUnknown,
		ProviderURL: r.app.providerName,
		ProviderOK:  false,
		Latency:     "",
		MCPCount:    0,
		MCPOnline:   0,
		GitOK:       false,
	}
	r.app.ShowDialog(screens.NewNetworkScreen(state))
	return "__dialog__"
}

func (r *commandRegistry) cmdHome(ctx context.Context, args []string) string {
	return "__home__"
}

func (r *commandRegistry) cmdRetry(ctx context.Context, args []string) string {
	if r.app.chatScreen == nil {
		return "Chat not initialized."
	}
	if r.app.state == StateCompleted || r.app.state == StateError || r.app.state == StateCancelled {
		r.app.state = StateIdle
		if len(r.app.history) > 0 {
			r.app.chatScreen.ClearToolCards()
			r.app.chatScreen.HideThinking()
			return r.app.startChat(ctx)
		}
		return "No previous response to retry."
	}
	return "Cannot retry in current state."
}

func (r *commandRegistry) cmdCancel(ctx context.Context, args []string) string {
	if r.app.chatScreen == nil {
		return "Chat not initialized."
	}
	if r.app.state != StateIdle && r.app.state != StateCompleted && r.app.state != StateCancelled {
		if r.app.cancel != nil {
			r.app.cancel()
			r.app.cancel = nil
		}
		r.app.state = StateCancelled
		r.app.chatScreen.HideThinking()
		r.app.chatScreen.ClearToolCards()
		r.app.partialContent = ""
		r.app.partialReasoning = ""
		r.app.statusBar.SetWorking(false)
		return "Operation cancelled."
	}
	return "Nothing to cancel."
}

func (r *commandRegistry) cmdTheme(ctx context.Context, args []string) string {
	themes := []string{"default", "dark", "light", "ocean", "forest", "monokai"}
	if len(args) == 0 {
		return "Usage: /theme <name>. Available: " + strings.Join(themes, ", ")
	}
	name := args[0]
	for _, t := range themes {
		if t == name {
			return "Theme set to: " + name
		}
	}
	return "Unknown theme. Available: " + strings.Join(themes, ", ")
}

func (r *commandRegistry) cmdPlugin(ctx context.Context, args []string) string {
	if r.app.pluginManager == nil {
		return "Plugin system not initialized."
	}
	if len(args) == 0 {
		plugins := r.app.pluginManager.List()
		if len(plugins) == 0 {
			return "No plugins loaded."
		}
		var b strings.Builder
		b.WriteString("Loaded plugins:\n")
		for _, p := range plugins {
			b.WriteString(fmt.Sprintf("  %s v%s\n", p.Name, p.Version))
		}
		return b.String()
	}
	switch args[0] {
	case "load":
		if len(args) < 2 {
			return "Usage: /plugin load <path>"
		}
		if err := r.app.pluginManager.Load(args[1]); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return "Plugin loaded."
	case "unload":
		if len(args) < 2 {
			return "Usage: /plugin unload <name>"
		}
		if err := r.app.pluginManager.Unload(args[1]); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return "Plugin unloaded."
	case "exec":
		if len(args) < 2 {
			return "Usage: /plugin exec <name> [args...]"
		}
		output, err := r.app.pluginManager.Execute(args[1], args[2:])
		if err != nil {
			return fmt.Sprintf("Plugin error: %v\nOutput: %s", err, output)
		}
		return output
	case "reload":
		if err := r.app.pluginManager.LoadAll(); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return "Plugins reloaded."
	default:
		return "Unknown subcommand. Available: load, unload, exec, reload"
	}
}

func (r *commandRegistry) cmdBatch(ctx context.Context, args []string) string {
	r.app.ShowDialog(screens.NewBatchEditScreen())
	return "__dialog__"
}

func (r *commandRegistry) cmdStop(ctx context.Context, args []string) string {
	if r.app.eventBus == nil {
		return "Event bus not initialized."
	}
	if r.app.state == StateIdle || r.app.state == StateCompleted || r.app.state == StateCancelled {
		return "No active generation to stop."
	}
	if r.app.cancel != nil {
		r.app.cancel()
		r.app.cancel = nil
	}
	r.app.state = StateCancelled
	r.app.statusBar.SetWorking(false)
	r.app.chatScreen.HideThinking()
	r.app.chatScreen.SetStreamActive(false)
	r.app.partialContent = r.app.streamBuf
	r.app.eventBus.Emit(eventbus.EventStreamComplete, nil)
	return "Generation stopped."
}

func (r *commandRegistry) cmdContinue(ctx context.Context, args []string) string {
	if r.app.chatScreen == nil {
		return "Chat not initialized."
	}
	if r.app.partialContent == "" && r.app.state != StatePaused {
		return "Nothing to continue from."
	}
	r.app.state = StateIdle
	if r.app.partialContent != "" {
		r.app.chatScreen.AddMessage("assistant", r.app.partialContent)
		r.app.saveMessage(session.RoleAssistant, r.app.partialContent, r.app.partialReasoning)
	}
	r.app.partialContent = ""
	r.app.partialReasoning = ""
	return "Continue your conversation below."
}

func (r *commandRegistry) cmdEdit(ctx context.Context, args []string) string {
	if len(args) < 2 {
		return "Usage: /edit <message_index> <new_content>"
	}
	idx := 0
	fmt.Sscanf(args[0], "%d", &idx)
	content := strings.Join(args[1:], " ")
	r.app.editMessage(idx-1, content)
	return fmt.Sprintf("Message %d edited.", idx)
}

func (r *commandRegistry) cmdDelete(ctx context.Context, args []string) string {
	if len(args) < 1 {
		return "Usage: /delete <message_index>"
	}
	idx := 0
	fmt.Sscanf(args[0], "%d", &idx)
	r.app.deleteMessage(idx - 1)
	return fmt.Sprintf("Message %d deleted.", idx)
}

func (r *commandRegistry) cmdRename(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return "Usage: /rename <new_session_name>"
	}
	name := strings.Join(args, " ")
	if r.app.currentSess != nil && r.app.sessionRepo != nil {
		r.app.currentSess.Name = name
		r.app.sessionRepo.Update(ctx, r.app.currentSess)
		return "Session renamed to: " + name
	}
	return "No active session to rename."
}

func (r *commandRegistry) cmdBranch(ctx context.Context, args []string) string {
	if len(args) == 0 {
		if len(r.app.branches) == 0 {
			return "No branches. Use /branch <name> to create one."
		}
		items := make([]screens.BranchItem, 0, len(r.app.branches))
		for i, b := range r.app.branches {
			items = append(items, screens.BranchItem{
				ID:       b.ID,
				Name:     b.Name,
				MsgCount: len(b.History),
				IsActive: i == r.app.currentBranch,
			})
		}
		bs := screens.NewBranchScreen()
		bs.SetBranches(items)
		r.app.ShowDialog(bs)
		return "__dialog__"
	}
	name := strings.Join(args, " ")
	r.app.createBranch(name)
	return "Branch created: " + name
}

func (r *commandRegistry) cmdUndo(ctx context.Context, args []string) string {
	r.app.undoLastEvent()
	return ""
}

func (r *commandRegistry) cmdRedo(ctx context.Context, args []string) string {
	r.app.redoLastEvent()
	return ""
}

func (r *commandRegistry) cmdExport(ctx context.Context, args []string) string {
	if r.app.currentSess == nil {
		return "No active session to export."
	}
	if r.app.sessionRepo == nil {
		return "Session storage not available."
	}
	data, err := exportSession(r.app.currentSess, r.app.history)
	if err != nil {
		return fmt.Sprintf("Export error: %v", err)
	}
	filename := fmt.Sprintf("tc-session-%s.json", r.app.currentSess.ID[:8])
	importFilePath := r.app.workspace + "/" + filename
	if len(args) > 0 && args[0] == "--path" && len(args) > 1 {
		importFilePath = strings.Join(args[1:], " ")
	}
	if err := saveToFile(importFilePath, data); err != nil {
		return fmt.Sprintf("Write error: %v", err)
	}
	return fmt.Sprintf("Session exported to %s", importFilePath)
}

func (r *commandRegistry) cmdImport(ctx context.Context, args []string) string {
	if r.app.messageRepo == nil {
		return "Message storage not available."
	}
	if len(args) == 0 {
		return "Usage: /import <filepath>"
	}
	path := strings.Join(args, " ")
	data, err := readFromFile(path)
	if err != nil {
		return fmt.Sprintf("Read error: %v", err)
	}
	sess, msgs, err := importSession(data)
	if err != nil {
		return fmt.Sprintf("Import error: %v", err)
	}
	r.app.currentSess = sess
	r.app.history = msgs
	r.app.chatScreen = screens.NewChatScreen()
	for _, msg := range msgs {
		r.app.chatScreen.AddMessage(string(msg.Role), msg.Content, msg.Reasoning)
	}
	r.app.screen = screenChat
	if r.app.sessionRepo != nil {
		r.app.sessionRepo.Create(ctx, sess)
		for i := range msgs {
			r.app.messageRepo.Create(ctx, &msgs[i])
		}
	}
	return fmt.Sprintf("Session '%s' imported (%d messages).", sess.Name, len(msgs))
}

func (r *commandRegistry) cmdPin(ctx context.Context, args []string) string {
	if r.app.currentSess == nil {
		return "No active session."
	}
	if r.app.sessionRepo != nil {
		r.app.currentSess.Status = session.StatusPinned
		r.app.sessionRepo.Update(ctx, r.app.currentSess)
		return "Session pinned."
	}
	return "Session storage not available."
}

func (r *commandRegistry) cmdCollab(ctx context.Context, args []string) string {
	if r.app.chatScreen == nil {
		return "Chat not initialized."
	}
	if len(args) == 0 {
		return "Usage: /collab start [port] | /collab connect <address> | /collab sync | /collab stop"
	}
	action := args[0]
	switch action {
	case "start":
		if r.app.collabServer != nil {
			return "Server already running."
		}
		port := "9876"
		if len(args) > 1 {
			port = args[1]
		}
		srv := collab.NewServer(port)
		r.app.collabServer = srv
		go func() {
			if err := srv.Start(); err != nil {
				r.app.eventBus.Emit(eventbus.EventError, "Collab server stopped: "+err.Error())
			}
		}()
		r.app.toastManager.Info("Collab server starting on port " + port)
		return "Collaboration server starting on port " + port + "."

	case "connect":
		if len(args) < 2 {
			return "Usage: /collab connect <address> (e.g., 192.168.1.100:9876)"
		}
		addr := args[1]
		if !strings.Contains(addr, "://") {
			addr = "http://" + addr
		}
		client := collab.NewClient(addr)
		sync, err := client.Get()
		if err != nil {
			return fmt.Sprintf("Connection failed: %v", err)
		}
		r.app.toastManager.Info("Connected to collab session: " + sync.SessionName)
		return fmt.Sprintf("Connected to collab session '%s' (%d messages). Use /collab sync to pull latest.",
			sync.SessionName, len(sync.Messages))

	case "sync":
		client := collab.NewClient("")
		sync, err := client.Get()
		if err != nil {
			return fmt.Sprintf("Sync failed: %v", err)
		}
		for _, msg := range sync.Messages {
			r.app.chatScreen.AddMessage(msg.Role, msg.Content, msg.Reasoning)
		}
		return fmt.Sprintf("Synced %d messages from collab session.", len(sync.Messages))

	case "push":
		if r.app.currentSess == nil {
			return "No active session to share."
		}
		sync := &collab.SessionSync{
			SessionID:   r.app.currentSess.ID,
			SessionName: r.app.currentSess.Name,
		}
		for _, msg := range r.app.history {
			sync.Messages = append(sync.Messages, collab.SyncMessage{
				Role:      string(msg.Role),
				Content:   msg.Content,
				Reasoning: msg.Reasoning,
				Timestamp: msg.CreatedAt,
			})
		}
		client := collab.NewClient("")
		if err := client.Push(sync); err != nil {
			return fmt.Sprintf("Push failed: %v", err)
		}
		return fmt.Sprintf("Session pushed to collab server (%d messages).", len(sync.Messages))

	case "stop":
		if r.app.collabServer != nil {
			r.app.collabServer.Stop()
			r.app.collabServer = nil
			return "Collaboration server stopped."
		}
		return "No collaboration server running."

	default:
		return "Usage: /collab start [port] | /collab connect <address> | /collab sync | /collab push | /collab stop"
	}
}

func (r *commandRegistry) cmdSearch(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return "Usage: /search <query>"
	}
	query := strings.ToLower(strings.Join(args, " "))
	if r.app.sessionRepo == nil {
		return "Session storage not available."
	}
	sessions, err := r.app.sessionRepo.ListActive(ctx)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	var results []screens.SessionListItem
	for _, s := range sessions {
		if strings.Contains(strings.ToLower(s.Name), query) {
			tokIn := s.TokenIn
			results = append(results, screens.SessionListItem{
				ID:        s.ID,
				Name:      s.Name,
				MsgCount:  s.MessageCnt,
				TokenIn:   tokIn,
				TokenOut:  s.TokenOut,
				IsActive:  r.app.currentSess != nil && s.ID == r.app.currentSess.ID,
				UpdatedAt: s.UpdatedAt,
			})
		}
	}
	if len(results) == 0 {
		return fmt.Sprintf("No sessions matching '%s'.", query)
	}
	ss := screens.NewSessionScreen()
	ss.SetSessions(results)
	r.app.ShowDialog(ss)
	return "__dialog__"
}

func exportSession(sess *session.Session, msgs []session.Message) ([]byte, error) {
	type exportData struct {
		Session  session.Session   `json:"session"`
		Messages []session.Message `json:"messages"`
	}
	return json.MarshalIndent(exportData{Session: *sess, Messages: msgs}, "", "  ")
}

func importSession(data []byte) (*session.Session, []session.Message, error) {
	type exportData struct {
		Session  session.Session   `json:"session"`
		Messages []session.Message `json:"messages"`
	}
	var ed exportData
	if err := json.Unmarshal(data, &ed); err != nil {
		return nil, nil, err
	}
	ed.Session.ID = uuid.New().String()
	for i := range ed.Messages {
		ed.Messages[i].ID = uuid.New().String()
		ed.Messages[i].SessionID = ed.Session.ID
	}
	return &ed.Session, ed.Messages, nil
}

func splitArgs(input string) []string {
	var args []string
	var current strings.Builder
	inQuote := false
	for _, r := range input {
		switch {
		case r == '"':
			inQuote = !inQuote
		case r == ' ' && !inQuote:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	return args
}

func saveToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func readFromFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
