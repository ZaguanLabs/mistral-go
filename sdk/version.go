package sdk

// Version information for the Mistral Go SDK
const (
	// Version is the current version of the SDK
	Version = "2.0.0"

	// SDKName is the name of the SDK
	SDKName = "mistral-go"

	// UserAgent is the user agent string sent with API requests
	UserAgent = SDKName + "/" + Version
)

// VersionInfo provides detailed version information
type VersionInfo struct {
	Version       string
	SDKName       string
	UserAgent     string
	GoVersion     string
	FeatureParity string
}

// GetVersionInfo returns detailed version information
func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:       Version,
		SDKName:       SDKName,
		UserAgent:     UserAgent,
		GoVersion:     "1.21+",
		FeatureParity: "100%",
	}
}
