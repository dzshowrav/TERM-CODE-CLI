package git

import (
	"fmt"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type Service struct{}

type StatusResult struct {
	Branch   string   `json:"branch"`
	Hash     string   `json:"hash"`
	Modified []string `json:"modified"`
	Added    []string `json:"added"`
	Deleted  []string `json:"deleted"`
	Staged   []string `json:"staged"`
	Ahead    int      `json:"ahead"`
	Behind   int      `json:"behind"`
	Clean    bool     `json:"clean"`
}

type LogEntry struct {
	Hash    string    `json:"hash"`
	Message string    `json:"message"`
	Author  string    `json:"author"`
	When    time.Time `json:"when"`
}

type DiffResult struct {
	Files []FileDiff `json:"files"`
}

type FileDiff struct {
	Name    string `json:"name"`
	Added   int    `json:"added"`
	Removed int    `json:"removed"`
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Open(path string) (*gogit.Repository, error) {
	repo, err := gogit.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("open repo: %w", err)
	}
	return repo, nil
}

func (s *Service) Status(repo *gogit.Repository) (*StatusResult, error) {
	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("worktree: %w", err)
	}

	status, err := wt.Status()
	if err != nil {
		return nil, fmt.Errorf("status: %w", err)
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("head: %w", err)
	}

	result := &StatusResult{
		Branch:   ref.Name().Short(),
		Hash:     ref.Hash().String()[:8],
		Modified: []string{},
		Added:    []string{},
		Deleted:  []string{},
		Staged:   []string{},
		Clean:    status.IsClean(),
	}

	for path, st := range status {
		if st.Staging != gogit.Untracked && st.Staging != gogit.Unmodified {
			result.Staged = append(result.Staged, path)
		}
		switch st.Worktree {
		case gogit.Modified:
			result.Modified = append(result.Modified, path)
		case gogit.Added:
			result.Added = append(result.Added, path)
		case gogit.Deleted:
			result.Deleted = append(result.Deleted, path)
		}
	}

	return result, nil
}

func (s *Service) Log(repo *gogit.Repository, count int) ([]LogEntry, error) {
	if count <= 0 {
		count = 10
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("head: %w", err)
	}

	iter, err := repo.Log(&gogit.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return nil, fmt.Errorf("log: %w", err)
	}

	var entries []LogEntry
	err = iter.ForEach(func(c *object.Commit) error {
		if len(entries) >= count {
			return fmt.Errorf("enough")
		}
		entries = append(entries, LogEntry{
			Hash:    c.Hash.String()[:8],
			Message: c.Message,
			Author:  c.Author.Name,
			When:    c.Author.When,
		})
		return nil
	})

	if err != nil && err.Error() != "enough" {
		return nil, err
	}

	return entries, nil
}

func (s *Service) Add(repo *gogit.Repository, files []string) error {
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("worktree: %w", err)
	}

	for _, file := range files {
		_, err := wt.Add(file)
		if err != nil {
			return fmt.Errorf("add %s: %w", file, err)
		}
	}
	return nil
}

func (s *Service) AddAll(repo *gogit.Repository) error {
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("worktree: %w", err)
	}

	_, err = wt.Add(".")
	return err
}

func (s *Service) Commit(repo *gogit.Repository, message string) (string, error) {
	wt, err := repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("worktree: %w", err)
	}

	hash, err := wt.Commit(message, &gogit.CommitOptions{
		Author: &object.Signature{
			Name:  "Term Code",
			Email: "tc@termcode.local",
			When:  time.Now(),
		},
	})
	if err != nil {
		return "", fmt.Errorf("commit: %w", err)
	}

	return hash.String()[:8], nil
}

func (s *Service) Diff(repo *gogit.Repository) (*DiffResult, error) {
	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("worktree: %w", err)
	}

	status, err := wt.Status()
	if err != nil {
		return nil, fmt.Errorf("status: %w", err)
	}

	result := &DiffResult{Files: []FileDiff{}}

	for path, st := range status {
		if st.Worktree == gogit.Unmodified && st.Staging == gogit.Unmodified {
			continue
		}
		fd := FileDiff{Name: path}
		if st.Worktree == gogit.Modified || st.Staging == gogit.Modified {
			fd.Added = 1
		}
		if st.Worktree == gogit.Deleted || st.Staging == gogit.Deleted {
			fd.Removed = 1
		}
		result.Files = append(result.Files, fd)
	}

	return result, nil
}

func (s *Service) Branches(repo *gogit.Repository) ([]string, error) {
	iter, err := repo.Branches()
	if err != nil {
		return nil, fmt.Errorf("branches: %w", err)
	}

	var branches []string
	err = iter.ForEach(func(ref *plumbing.Reference) error {
		branches = append(branches, ref.Name().Short())
		return nil
	})
	return branches, err
}

func (s *Service) Checkout(repo *gogit.Repository, branch string, create bool) error {
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("worktree: %w", err)
	}

	opts := &gogit.CheckoutOptions{
		Create: create,
		Force:  false,
	}

	if create {
		opts.Branch = plumbing.NewBranchReferenceName(branch)
	} else {
		opts.Branch = plumbing.NewBranchReferenceName(branch)
	}

	return wt.Checkout(opts)
}

func (s *Service) IsRepo(path string) bool {
	_, err := gogit.PlainOpen(path)
	return err == nil
}

func (s *Service) GetBranch(repo *gogit.Repository) (string, error) {
	ref, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("head: %w", err)
	}
	return ref.Name().Short(), nil
}

var _ = transport.UnsupportedCapabilities
