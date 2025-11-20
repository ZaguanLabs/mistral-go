package sdk

import (
	"strings"
	"testing"
)

func TestVersionConstants(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}

	if SDKName == "" {
		t.Error("SDKName should not be empty")
	}

	if UserAgent == "" {
		t.Error("UserAgent should not be empty")
	}
}

func TestVersionFormat(t *testing.T) {
	// Version should be in semver format (e.g., "2.0.0")
	parts := strings.Split(Version, ".")
	if len(parts) != 3 {
		t.Errorf("Version should be in semver format (x.y.z), got %s", Version)
	}
}

func TestSDKName(t *testing.T) {
	if SDKName != "mistral-go" {
		t.Errorf("Expected SDKName to be 'mistral-go', got %s", SDKName)
	}
}

func TestUserAgentFormat(t *testing.T) {
	expected := SDKName + "/" + Version
	if UserAgent != expected {
		t.Errorf("Expected UserAgent to be '%s', got '%s'", expected, UserAgent)
	}
}

func TestGetVersionInfo(t *testing.T) {
	info := GetVersionInfo()

	if info.Version != Version {
		t.Errorf("Expected Version %s, got %s", Version, info.Version)
	}

	if info.SDKName != SDKName {
		t.Errorf("Expected SDKName %s, got %s", SDKName, info.SDKName)
	}

	if info.UserAgent != UserAgent {
		t.Errorf("Expected UserAgent %s, got %s", UserAgent, info.UserAgent)
	}

	if info.GoVersion == "" {
		t.Error("GoVersion should not be empty")
	}

	if info.FeatureParity == "" {
		t.Error("FeatureParity should not be empty")
	}
}

func TestVersionInfoStructure(t *testing.T) {
	info := GetVersionInfo()

	// Check that all fields are populated
	if info.Version == "" || info.SDKName == "" || info.UserAgent == "" ||
		info.GoVersion == "" || info.FeatureParity == "" {
		t.Error("All VersionInfo fields should be populated")
	}
}

func TestFeatureParity(t *testing.T) {
	info := GetVersionInfo()

	if info.FeatureParity != "100%" {
		t.Errorf("Expected FeatureParity to be '100%%', got %s", info.FeatureParity)
	}
}
