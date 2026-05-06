package fixtures

import (
	"github.com/Masterminds/semver/v3"
)

const (
	Version_V1_0_0 = "v1.0.0"
	Version_V1_0_1 = "v1.0.1"
)

var (
	SemVer_V1_0_0 = semver.MustParse(Version_V1_0_0)
	SemVer_V1_0_1 = semver.MustParse(Version_V1_0_1)
)
