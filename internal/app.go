package internal

import "github.com/praveenmahasena/goImg/internal/files"

func Start() error {
	_, pathErr := files.PathExists()

	if pathErr != nil {
		return pathErr
	}

	return nil
}
