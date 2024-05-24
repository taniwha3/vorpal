package build

import (
	"github.com/erikreinert/vorpal/internal/artifact"
	"github.com/erikreinert/vorpal/internal/store"
)

func BuildArtifact(a artifact.Artifact) error {
	artifactDir := store.GetStorePath(a.Name)

	err := store.NewStoreDir(artifactDir)
	if err != nil {
		return err
	}

	err = store.CopyDir(a.Source, artifactDir, a.Ignore)
	if err != nil {
		return err
	}

	return nil
}
