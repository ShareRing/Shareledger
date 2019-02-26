package version

// Version components
const (
	Maj = "0"
	Min = "1"
	Fix = "0"
)

var (
	Version = "0.2.0"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
