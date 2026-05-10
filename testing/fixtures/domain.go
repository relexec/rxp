package fixtures

import "github.com/relexec/rxp/domain"

const (
	DomainName = "domain.testing.fixtures.rxp"
)

var (
	Domain = domain.New(
		domain.WithName(DomainName),
	)
)
