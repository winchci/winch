// Code generated by winch. DO NOT EDIT.
package version

import (
	"fmt"
	"runtime"
)

const (
	Name        = "winch"
	Description = "Universal build and release tool"
	ReleaseName = "bright pronghorn"
	Version     = "0.26.5"
	Prerelease  = ""
)

// String returns the complete version string, including prerelease
func String() string {
	s := fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s %s", Version, Prerelease, s)
	}
	return fmt.Sprintf("%s %s", Version, s)
}
