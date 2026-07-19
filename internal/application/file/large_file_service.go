package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	DefaultChunkSize = 64 * 1024
	DefaultMaxSize   = 100 * 1024 * 1024
	WarnSize         = 10 * 1024 * 1024
	PreviewLines     = 100
)

type LargeFileService struct {
	MaxSize   int
	ChunkSize int
}

func NewLargeFileService() *LargeFileService {
	return &LargeFileService{
		MaxSize:   DefaultMaxSize,
		ChunkSize: DefaultChunkSize,
	}
}

type FileInfo struct {
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	IsLarge   bool   `json:"is_large"`
	WarnSize  bool   `json:"warn_size"`
	LineCount int    `json:"line_count,omitempty"`
}

func (s *LargeFileService) Check(path string) (*FileInfo, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	fi := &FileInfo{
		Path:     abs,
		Size:     info.Size(),
		IsLarge:  info.Size() > int64(s.MaxSize),
		WarnSize: info.Size() > int64(WarnSize),
	}

	if info.Size() > 0 {
		f, err := os.Open(abs)
		if err != nil {
			return fi, nil
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fi.LineCount++
		}
	}

	return fi, nil
}

func (s *LargeFileService) ReadHead(path string, maxLines int) ([]string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	f, err := os.Open(abs)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, s.ChunkSize), s.ChunkSize)
	for scanner.Scan() && len(lines) < maxLines {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func (s *LargeFileService) ReadTail(path string, maxLines int) ([]string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	f, err := os.Open(abs)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, s.ChunkSize), s.ChunkSize)

	// Read into ring buffer
	ring := make([]string, maxLines)
	idx := 0
	count := 0

	for scanner.Scan() {
		ring[idx%maxLines] = scanner.Text()
		idx++
		count++
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	if count < maxLines {
		return ring[:count], nil
	}

	// Rotate ring buffer
	result := make([]string, maxLines)
	for i := 0; i < maxLines; i++ {
		result[i] = ring[(idx)%maxLines]
		idx++
	}
	return result, nil
}

func (s *LargeFileService) ReadRange(path string, startLine, endLine int) ([]string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	f, err := os.Open(abs)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, s.ChunkSize), s.ChunkSize)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		if lineNum >= startLine && lineNum <= endLine {
			lines = append(lines, scanner.Text())
		}
		if lineNum > endLine {
			break
		}
	}

	return lines, scanner.Err()
}

func (s *LargeFileService) StreamWrite(path string, reader io.Reader) (int64, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return 0, fmt.Errorf("resolve path: %w", err)
	}

	dir := filepath.Dir(abs)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return 0, fmt.Errorf("mkdir: %w", err)
	}

	f, err := os.Create(abs)
	if err != nil {
		return 0, fmt.Errorf("create: %w", err)
	}
	defer f.Close()

	return io.Copy(f, reader)
}

func (s *LargeFileService) ReadJSON(path string, v any) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}

	info, err := os.Stat(abs)
	if err != nil {
		return err
	}

	if info.Size() > int64(s.MaxSize) {
		return fmt.Errorf("file too large: %d bytes (max %d)", info.Size(), s.MaxSize)
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (s *LargeFileService) ReadStream(path string, fn func(reader *bufio.Scanner) error) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}

	f, err := os.Open(abs)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, s.ChunkSize), s.ChunkSize)
	return fn(scanner)
}

func (s *LargeFileService) CopyStream(src, dst string) (int64, error) {
	absSrc, err := filepath.Abs(src)
	if err != nil {
		return 0, fmt.Errorf("resolve src: %w", err)
	}
	absDst, err := filepath.Abs(dst)
	if err != nil {
		return 0, fmt.Errorf("resolve dst: %w", err)
	}

	dir := filepath.Dir(absDst)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return 0, fmt.Errorf("mkdir: %w", err)
	}

	r, err := os.Open(absSrc)
	if err != nil {
		return 0, fmt.Errorf("open src: %w", err)
	}
	defer r.Close()

	w, err := os.Create(absDst)
	if err != nil {
		return 0, fmt.Errorf("create dst: %w", err)
	}
	defer w.Close()

	return io.Copy(w, r)
}
