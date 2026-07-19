package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type MoveService struct{}

func NewMoveService() *MoveService {
	return &MoveService{}
}

func (s *MoveService) Move(source, dest string) error {
	absSrc, err := filepath.Abs(source)
	if err != nil {
		return fmt.Errorf("resolve source: %w", err)
	}
	absDst, err := filepath.Abs(dest)
	if err != nil {
		return fmt.Errorf("resolve dest: %w", err)
	}

	dir := filepath.Dir(absDst)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create dest dir: %w", err)
	}

	if err := os.Rename(absSrc, absDst); err != nil {
		return fmt.Errorf("rename: %w", err)
	}
	return nil
}

type CopyService struct{}

func NewCopyService() *CopyService {
	return &CopyService{}
}

func (s *CopyService) Copy(source, dest string) error {
	absSrc, err := filepath.Abs(source)
	if err != nil {
		return fmt.Errorf("resolve source: %w", err)
	}
	absDst, err := filepath.Abs(dest)
	if err != nil {
		return fmt.Errorf("resolve dest: %w", err)
	}

	srcInfo, err := os.Stat(absSrc)
	if err != nil {
		return fmt.Errorf("stat source: %w", err)
	}

	dir := filepath.Dir(absDst)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create dest dir: %w", err)
	}

	if srcInfo.IsDir() {
		return s.copyDir(absSrc, absDst)
	}
	return s.copyFile(absSrc, absDst, srcInfo.Mode())
}

func (s *CopyService) copyFile(src, dst string, mode os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open source: %w", err)
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL|os.O_WRONLY, mode)
	if err != nil {
		return fmt.Errorf("create dest: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy data: %w", err)
	}
	return nil
}

func (s *CopyService) copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return fmt.Errorf("mkdir dest: %w", err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("read source dir: %w", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("info %s: %w", entry.Name(), err)
		}

		if info.IsDir() {
			if err := s.copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := s.copyFile(srcPath, dstPath, info.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}
