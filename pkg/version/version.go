package version

import (
	"encoding/json"
	"os"
	"runtime"
	"time"
)

var (
	Env       string
	Repo      string
	Version   string
	Swagger   string
	Build     string
	BuildTime string
	Uptime    string
	CreatedAt time.Time
)

// Build version info.
type BuildInfo struct {
	GoVersion string
	Env       string
	Repo      string
	Version   string
	Swagger   string `json:"swagger,omitempty"`
	Build     string
	BuildTime string
	Uptime    string
	CreatedAt time.Time
}

func init() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	CreatedAt = time.Now().In(loc)

	Env = os.Getenv("ENV")
}

func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		GoVersion: runtime.Version(),
		Env:       Env,
		Repo:      Repo,
		Version:   Version,
		Swagger:   Swagger,
		Build:     Build,
		BuildTime: BuildTime,
		CreatedAt: CreatedAt,
		Uptime:    time.Since(CreatedAt).String(),
	}
}

func (info *BuildInfo) ToString() string {
	out, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	return string(out)
}
