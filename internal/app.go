package internal

import (
	"github.com/praveenmahasena/goImg/internal/files"
)

func Start() error {
	return files.ReadDir()
}
