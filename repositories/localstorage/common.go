package localstorage

import (
	"fmt"
	"os"
	"path/filepath"
)

const WriteTestFile = "write_test.tmp"

type LocalStorage struct {
	root string
}

func NewLocalStorage(rootPath string) (*LocalStorage, error) {
	// Check if path exists => If yes, should be a folder
	rootAbs, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}
	folderInfo, err := os.Stat(rootAbs)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if err == nil && !folderInfo.IsDir() {
		return nil, fmt.Errorf("provided rootPath points to an existing file: %s", rootAbs)
	}

	// Folder doesn't exist => Create
	err = os.MkdirAll(rootAbs, 0o755)
	if err != nil {
		return nil, err
	}

	// Build repo
	repo := &LocalStorage{root: rootPath}

	// Check if folder is writable
	err = repo.SaveFile(WriteTestFile, []byte("test"))
	if err != nil {
		return nil, err
	}
	err = repo.DeleteFile(WriteTestFile)
	if err != nil {
		return nil, err
	}

	// New repo successful
	return repo, nil
}

func (s *LocalStorage) pathFromName(filename string) string {
	return filepath.Join(s.root, filename)
}
