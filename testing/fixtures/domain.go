package fixtures

import "github.com/relexec/rxp/domain"

const (
	DomainName = "testing.fixtures.rxp"
)

var (
	Domain = domain.New(
		domain.WithName(DomainName),
	)
)
