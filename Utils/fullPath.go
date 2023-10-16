package Utils

import (
	"os"
	"path/filepath"
)

func FullPath(relativePath string) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := exPath + "/" + relativePath
	return fullPath
}
