package fixtures

import "github.com/relexec/rxp/system"

const (
	SystemUUID = "9477c2b9-0a90-4546-8364-148e4a6dde08"
	SystemName = "testing.fixtures.rxp"
)

var (
	System = system.New(
		system.WithUUID(SystemUUID),
		system.WithName(SystemName),
	)
)
