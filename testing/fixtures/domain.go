package fixtures

import "github.com/relexec/rxp/domain"

const (
	DomainUUID = "c33d0da3-dea5-4bb3-854f-3bbe0d1ee959"
	DomainName = "domain.testing.fixtures.rxp"
)

var (
	Domain = domain.New(
		domain.WithUUID(DomainUUID),
		domain.WithName(DomainName),
	)
)
