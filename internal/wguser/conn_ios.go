package wguser

import (
	"os"
	"strings"
)

func init() {
	// in iOS there is no access to write into "/var/run" or "/var/run/wireguard" directory
	// but we can use temp directory under iOS
	socketDirectory = strings.TrimSuffix(os.TempDir(), string(os.PathSeparator))
}
