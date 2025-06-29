package version

const (
	Version = "0.1.0"

	BuildDate = "2025-06-26"

	CommitHash = "dev"
)

func GetVersion() string {
	return Version
}

func GetBuildInfo() map[string]string {
	return map[string]string{
		"version":    Version,
		"buildDate":  BuildDate,
		"commitHash": CommitHash,
	}
}
