package artifact

type (
	Ignore []string
	Name   string
	Source string
)

type Artifact struct {
	Ignore Ignore
	Name   Name
	Source Source
}
