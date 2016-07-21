// Allows the version string to be set at build time
// and provides an exported method for external access

package config

// This is defined at build time via build flags: -ldflags "-X {{ cookiecutter.project_name|lower }}/config.version=abcdefg"
var version string

// Exported method for returning the version string
func Version() string {
	if version == "" {
		return "unknown"
	}
	return version
}
