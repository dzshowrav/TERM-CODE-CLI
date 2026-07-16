package tui

import (
	"context"
	"fmt"
	"strings"

	"termcode/internal/adapters/tui/screens"
	"termcode/internal/application/model"
	"termcode/internal/application/provider"
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
	r.register("model", func(ctx context.Context, args []string) string {
		return r.cmdModel(ctx, args, providerSvc, modelSvc)
	})
	r.register("agent", r.cmdAgent)
	r.register("workspace", r.cmdWorkspace)
	r.register("sessions", r.cmdSessions)
	r.register("session", r.cmdSessions)
	r.register("clear", r.cmdClear)
	r.register("git", r.cmdGit)
	r.register("exit", r.cmdExit)
	r.register("quit", r.cmdExit)
	r.register("home", r.cmdHome)

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

	parts := strings.Fields(input[1:])
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
	return `Commands:
  /help, /h           Show this help
  /provider list      List providers
  /provider add       Add a provider
  /provider select    Select active provider
  /provider sync      Sync models from provider
  /model list         List models
  /model select       Select active model
  /agent list         List agents
  /agent select       Select active agent
  /workspace <path>   Set workspace path
  /sessions list      List saved sessions
  /sessions new       Start a new session
  /sessions delete    Delete a session
  /git status         Show git status
  /git log [n]        Show commit log (last n commits)
  /git diff           Show working tree changes
  /git add [files]    Stage files
  /git commit <msg>   Commit staged changes
  /git branches       List branches
  /clear              Clear session / return home
  /home               Return to home screen
  /exit, /quit        Close TermCode`
}

func (r *commandRegistry) cmdProvider(ctx context.Context, args []string, providerSvc *provider.Service, modelSvc *model.Service) string {
	if len(args) == 0 {
		return "Usage: /provider list|add|select|delete|sync"
	}

	switch args[0] {
	case "list":
		providers, err := providerSvc.List(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if len(providers) == 0 {
			return "No providers configured. Use /provider add to add one."
		}
		var result []string
		for _, p := range providers {
			def := ""
			if p.IsDefault {
				def = " (active)"
			}
			result = append(result, fmt.Sprintf("  %s - %s%s", p.Name, p.BaseURL, def))
		}
		return "Providers:\n" + strings.Join(result, "\n")

	case "add":
		if len(args) < 3 {
			return "Usage: /provider add <name> <base_url> [api_key]"
		}
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

	case "select":
		if len(args) < 2 {
			return "Usage: /provider select <name>"
		}
		providers, err := providerSvc.List(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		name := args[1]
		for _, p := range providers {
			if strings.EqualFold(p.Name, name) {
				if err := providerSvc.SetDefault(ctx, p.ID); err != nil {
					return fmt.Sprintf("Error: %v", err)
				}
				r.app.providerName = p.Name
				r.app.homeScreen.UpdateConfig(screens.HomeScreenConfig{ProviderName: p.Name, ProviderURL: p.BaseURL})
				r.app.statusBar.SetModel(p.Name + "/" + r.app.modelName)
				return fmt.Sprintf("Provider '%s' is now active.", p.Name)
			}
		}
		return fmt.Sprintf("Provider '%s' not found.", name)

	case "delete":
		if len(args) < 2 {
			return "Usage: /provider delete <name>"
		}
		providers, err := providerSvc.List(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		name := args[1]
		for _, p := range providers {
			if strings.EqualFold(p.Name, name) {
				if err := providerSvc.Delete(ctx, p.ID); err != nil {
					return fmt.Sprintf("Error: %v", err)
				}
				return fmt.Sprintf("Provider '%s' deleted.", p.Name)
			}
		}
		return fmt.Sprintf("Provider '%s' not found.", name)

	case "sync":
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

	default:
		return "Usage: /provider list|add|select|delete|sync"
	}
}

func (r *commandRegistry) cmdModel(ctx context.Context, args []string, providerSvc *provider.Service, modelSvc *model.Service) string {
	if len(args) == 0 {
		return "Usage: /model list|select"
	}

	switch args[0] {
	case "list":
		models, err := modelSvc.List(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if len(models) == 0 {
			return "No models configured. Use /provider sync to load models."
		}
		var result []string
		for _, m := range models {
			fav := ""
			if m.IsFavorite {
				fav = " *"
			}
			result = append(result, fmt.Sprintf("  %s - %s (%s)%s", m.DisplayName, string(m.Category), m.ModelID, fav))
		}
		return "Models:\n" + strings.Join(result, "\n")

	case "select":
		if len(args) < 2 {
			return "Usage: /model select <model_id>"
		}
		models, err := modelSvc.List(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		q := args[1]
		for _, m := range models {
			if strings.EqualFold(m.ModelID, q) || strings.EqualFold(m.DisplayName, q) {
				r.app.SetSelectedModel(m.ModelID)
				return fmt.Sprintf("Model '%s' selected.", m.DisplayName)
			}
		}
		return fmt.Sprintf("Model '%s' not found. Use /model list to see available models.", q)

	default:
		return "Usage: /model list|select"
	}
}

func (r *commandRegistry) cmdAgent(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return "Usage: /agent list|select"
	}

	switch args[0] {
	case "list":
		return `Agents:
  General     - General-purpose coding assistant
  Expert      - Expert-level code review and debugging
  Architect   - System design and architecture`

	case "select":
		if len(args) < 2 {
			return "Usage: /agent select <name>"
		}
		name := args[1]
		r.app.agentName = name
		r.app.homeScreen.UpdateConfig(screens.HomeScreenConfig{AgentName: name})
		r.app.statusBar.SetAgent(name)
		return fmt.Sprintf("Agent '%s' selected.", name)

	default:
		return "Usage: /agent list|select"
	}
}

func (r *commandRegistry) cmdWorkspace(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return fmt.Sprintf("Current workspace: %s", r.app.workspace)
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

	if len(args) == 0 {
		return "Usage: /sessions list|new|delete"
	}

	switch args[0] {
	case "list":
		sessions, err := r.app.sessionRepo.ListActive(ctx)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if len(sessions) == 0 {
			return "No saved sessions."
		}
		var result []string
		for _, s := range sessions {
			msg := fmt.Sprintf("  %s (%d msgs)", s.Name, s.MessageCnt)
			if r.app.currentSess != nil && s.ID == r.app.currentSess.ID {
				msg += " [active]"
			}
			result = append(result, msg)
		}
		return "Sessions:\n" + strings.Join(result, "\n")

	case "new":
		r.app.history = nil
		r.app.currentSess = nil
		r.app.screen = screenHome
		r.app.homeScreen.UpdateConfig(screens.HomeScreenConfig{})
		return "New session started."

	case "delete":
		if len(args) < 2 {
			return "Usage: /sessions delete <session_id>"
		}
		id := args[1]
		if r.app.messageRepo != nil {
			r.app.messageRepo.DeleteBySession(ctx, id)
		}
		if err := r.app.sessionRepo.Delete(ctx, id); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return fmt.Sprintf("Session '%s' deleted.", id)

	default:
		return "Usage: /sessions list|new|delete"
	}
}

func (r *commandRegistry) cmdGit(ctx context.Context, args []string) string {
	if len(args) == 0 {
		return "Usage: /git status|log|diff|add|commit|branches"
	}

	switch args[0] {
	case "status":
		svc := git.NewService()
		if !svc.IsRepo(r.app.workspace) {
			return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
		}
		repo, err := svc.Open(r.app.workspace)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		status, err := svc.Status(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if status.Clean {
			return fmt.Sprintf("On branch %s. Clean working tree.", status.Branch)
		}
		var parts []string
		parts = append(parts, fmt.Sprintf("On branch %s (%s)", status.Branch, status.Hash))
		if len(status.Staged) > 0 {
			parts = append(parts, "Staged: "+strings.Join(status.Staged, ", "))
		}
		if len(status.Modified) > 0 {
			parts = append(parts, "Modified: "+strings.Join(status.Modified, ", "))
		}
		if len(status.Added) > 0 {
			parts = append(parts, "Added: "+strings.Join(status.Added, ", "))
		}
		if len(status.Deleted) > 0 {
			parts = append(parts, "Deleted: "+strings.Join(status.Deleted, ", "))
		}
		return strings.Join(parts, "\n")

	case "log":
		svc := git.NewService()
		if !svc.IsRepo(r.app.workspace) {
			return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
		}
		repo, err := svc.Open(r.app.workspace)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
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
			lines = append(lines, fmt.Sprintf("%s %s - %s", e.Hash, e.When.Format("2006-01-02"), msg))
		}
		return strings.Join(lines, "\n")

	case "diff":
		svc := git.NewService()
		if !svc.IsRepo(r.app.workspace) {
			return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
		}
		repo, err := svc.Open(r.app.workspace)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		diff, err := svc.Diff(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if len(diff.Files) == 0 {
			return "No changes."
		}
		var lines []string
		for _, f := range diff.Files {
			lines = append(lines, fmt.Sprintf("%s (+%d/-%d)", f.Name, f.Added, f.Removed))
		}
		return strings.Join(lines, "\n")

	case "add":
		svc := git.NewService()
		if !svc.IsRepo(r.app.workspace) {
			return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
		}
		repo, err := svc.Open(r.app.workspace)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		if len(args) > 1 {
			files := args[1:]
			if err := svc.Add(repo, files); err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Staged: %s", strings.Join(files, ", "))
		}
		if err := svc.AddAll(repo); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return "All files staged."

	case "commit":
		if len(args) < 2 {
			return "Usage: /git commit <message>"
		}
		svc := git.NewService()
		if !svc.IsRepo(r.app.workspace) {
			return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
		}
		repo, err := svc.Open(r.app.workspace)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		message := strings.Join(args[1:], " ")
		hash, err := svc.Commit(repo, message)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return fmt.Sprintf("Committed as %s: %s", hash, message)

	case "branches":
		svc := git.NewService()
		if !svc.IsRepo(r.app.workspace) {
			return fmt.Sprintf("'%s' is not a git repository.", r.app.workspace)
		}
		repo, err := svc.Open(r.app.workspace)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		branches, err := svc.Branches(repo)
		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}
		return "Branches:\n  " + strings.Join(branches, "\n  ")

	default:
		return "Usage: /git status|log|diff|add|commit|branches"
	}
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

func (r *commandRegistry) cmdHome(ctx context.Context, args []string) string {
	return "__home__"
}
