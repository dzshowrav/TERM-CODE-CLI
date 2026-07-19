package docs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type DocEntry struct {
	Path      string    `json:"path"`
	Title     string    `json:"title"`
	Content   string    `json:"content,omitempty"`
	Tags      []string  `json:"tags"`
	IndexedAt time.Time `json:"indexed_at"`
}

type Indexer struct {
	mu      sync.RWMutex
	entries map[string]DocEntry
	rootDir string
	watcher *fsnotify.Watcher
	stopCh  chan struct{}
}

func NewIndexer(rootDir string) *Indexer {
	return &Indexer{
		entries: make(map[string]DocEntry),
		rootDir: rootDir,
		stopCh:  make(chan struct{}),
	}
}

func (idx *Indexer) Start() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %w", err)
	}
	idx.watcher = watcher

	if err := idx.watchRecursive(idx.rootDir); err != nil {
		watcher.Close()
		return fmt.Errorf("watch recursive: %w", err)
	}

	idx.IndexAll()

	go idx.watchLoop()

	return nil
}

func (idx *Indexer) Stop() error {
	close(idx.stopCh)
	if idx.watcher != nil {
		return idx.watcher.Close()
	}
	return nil
}

func (idx *Indexer) watchRecursive(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return idx.watcher.Add(path)
		}
		return nil
	})
}

func (idx *Indexer) watchLoop() {
	for {
		select {
		case event, ok := <-idx.watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Create|fsnotify.Write) != 0 {
				info, err := os.Stat(event.Name)
				if err == nil && !info.IsDir() {
					idx.IndexFile(event.Name)
				}
			}
		case err, ok := <-idx.watcher.Errors:
			if !ok {
				return
			}
			_ = err
		case <-idx.stopCh:
			return
		}
	}
}

func (idx *Indexer) IndexFile(path string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	relPath, _ := filepath.Rel(idx.rootDir, path)
	content := string(data)

	entry := DocEntry{
		Path:      relPath,
		Title:     guessTitle(content, relPath),
		Content:   content,
		Tags:      extractTags(content),
		IndexedAt: time.Now(),
	}
	idx.entries[relPath] = entry
}

func (idx *Indexer) IndexAll() {
	filepath.Walk(idx.rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			idx.IndexFile(path)
		}
		return nil
	})
}

func (idx *Indexer) Search(query string) []DocEntry {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	query = strings.ToLower(query)
	var results []DocEntry
	for _, entry := range idx.entries {
		if strings.Contains(strings.ToLower(entry.Title), query) ||
			strings.Contains(strings.ToLower(entry.Content), query) {
			results = append(results, entry)
		}
	}
	return results
}

func (idx *Indexer) Get(path string) (DocEntry, bool) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	entry, ok := idx.entries[path]
	return entry, ok
}

func (idx *Indexer) Count() int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return len(idx.entries)
}

func guessTitle(content, path string) string {
	lines := strings.SplitN(content, "\n", 5)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
		if strings.HasPrefix(line, "// ") {
			return strings.TrimPrefix(line, "// ")
		}
	}
	return filepath.Base(path)
}

func extractTags(content string) []string {
	var tags []string
	seen := make(map[string]bool)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "// ") {
			text := strings.TrimPrefix(line, "// ")
			if strings.HasPrefix(text, "Package ") {
				tag := strings.Fields(text)[1]
				if !seen[tag] {
					tags = append(tags, tag)
					seen[tag] = true
				}
			}
		}
	}

	return tags
}
