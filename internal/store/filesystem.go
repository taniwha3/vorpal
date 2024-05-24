package store

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/erikreinert/vorpal/internal/artifact"
)

const Store = "/tmp/vorpal/store"

type Path string

func GetStorePath(name artifact.Name) Path {
	return Path(Store + "/" + name)
}

func NewStoreDir(path Path) error {
	err := os.MkdirAll(string(path), 0o755)
	if err != nil {
		return err
	}
	return nil
}

func CopyDir(src artifact.Source, dst Path, ignore artifact.Ignore) error {
	source := string(src)

	destination := string(dst)

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, ignorePath := range ignore {
			if strings.Contains(path, ignorePath) {
				return nil
			}
		}

		destPath := filepath.Join(destination, strings.TrimPrefix(path, source))
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}

		defer destFile.Close()

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}

		defer srcFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}

		log.Printf("Copied %s to %s", path, destPath)

		return os.Chmod(destPath, info.Mode())
	})
}
