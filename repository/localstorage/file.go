package localstorage

import (
	"os"
)

func (s *LocalStorage) SaveFile(filename string, content []byte) error {
	return os.WriteFile(s.pathFromName(filename), content, 0444)
}

func (s *LocalStorage) DeleteFile(filename string) error {
	return os.Remove(s.pathFromName(filename))
}
