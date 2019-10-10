package task_ffmpeg

import (
	"fmt"
	"runtime"
)

var (
	// specifiy version, BuildTimeUTC, AppName at build time with `-ldflags "-X path.to.package.Version x.x.x"` etc...
	Version      = "-"
	BuildTimeUTC = "-"
)

func GetVersion() string {
	return fmt.Sprintf(
		"%s (built w/%s)\nUTC Build Time: %v",
		Version,
		runtime.Version(),
		BuildTimeUTC,
	)
}
