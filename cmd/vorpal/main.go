package main

import (
	"log"

	"github.com/erikreinert/vorpal/internal/artifact"
	"github.com/erikreinert/vorpal/internal/build"
)

func main() {
	example := artifact.Artifact{
		Ignore: artifact.Ignore{".git", ".gitignore", ".direnv"},
		Name:   artifact.Name("example"),
		Source: artifact.Source("."),
	}

	err := build.BuildArtifact(example)
	if err != nil {
		log.Fatal(err)
	}
}
