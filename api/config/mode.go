package config

// Running modes
const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

var (
	modeName = DebugMode
)

func init() {
	if config.Debug {
		SetMode("debug")
	} else {
		SetMode("release")
	}
}

// SetMode sets system running mode, e.g. config.DebugMode.
func SetMode(value string) {
	switch value {
	case DebugMode:
		modeName = DebugMode
	case ReleaseMode, "":
		modeName = ReleaseMode
	default:
		panic("system running mode unknown: " + value)
	}
}

// Mode returns current running mode.
func Mode() string {
	return modeName
}

// IsDebugMode tells if running in debug mode.
func IsDebugMode() bool {
	return modeName == DebugMode
}
