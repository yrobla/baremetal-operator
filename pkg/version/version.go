// Package version includes the version information.
package version

import "fmt"

var (
	// Raw is the string representation of the version. This will be replaced
	// with the calculated version at build time.
	// set in hack/build.sh
	Raw = "was not built with version info"

	// String is the human-friendly representation of the version.
	String = fmt.Sprintf("metal3-io/baremetal-operator %s", Raw)

	// Commit is the commit hash from which the installer was built.
	// Set via LDFLAGS in Makefile.
	Commit = ""
)
