package fixtures

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
)

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

// A set of domains to test the hierarchical domain functionality
var (
	DomainTree_RootUUID = "b50aae84-1d51-49c0-9f4e-9aad727c3b51"
	DomainTree_RootName = api.DomainName("tree.testing.fixtures.rxp")
	DomainTree_Root     = domain.New(
		domain.WithUUID(DomainTree_RootUUID),
		domain.WithName(DomainTree_RootName),
	)
	DomainTree_Group1UUID = "6ce2cc0c-50f3-4b6e-a0e8-2f7218e93468"
	DomainTree_Group1Name = api.DomainName("group1.tree.testing.fixtures.rxp")
	DomainTree_Group1     = domain.New(
		domain.WithUUID(DomainTree_Group1UUID),
		domain.WithName(DomainTree_Group1Name),
		domain.WithParent(DomainTree_Root),
	)
	DomainTree_Group1Leaf1UUID = "61f52f2a-dd60-42a0-ae5f-1b07f52374bc"
	DomainTree_Group1Leaf1Name = api.DomainName("leaf1.group1.tree.testing.fixtures.rxp")
	DomainTree_Group1Leaf1     = domain.New(
		domain.WithUUID(DomainTree_Group1Leaf1UUID),
		domain.WithName(DomainTree_Group1Leaf1Name),
		domain.WithParent(DomainTree_Group1),
	)
	DomainTree_Group1Leaf2UUID = "6f6a202b-c040-4696-9867-4d51555c17aa"
	DomainTree_Group1Leaf2Name = api.DomainName("leaf2.group1.tree.testing.fixtures.rxp")
	DomainTree_Group1Leaf2     = domain.New(
		domain.WithUUID(DomainTree_Group1Leaf2UUID),
		domain.WithName(DomainTree_Group1Leaf2Name),
		domain.WithParent(DomainTree_Group1),
	)
	DomainTree_Group2UUID = "fcc59ed2-0faa-4004-a2af-441492625440"
	DomainTree_Group2Name = api.DomainName("group2.tree.testing.fixtures.rxp")
	DomainTree_Group2     = domain.New(
		domain.WithUUID(DomainTree_Group2UUID),
		domain.WithName(DomainTree_Group2Name),
		domain.WithParent(DomainTree_Root),
	)
	DomainTree_Group2Leaf1UUID = "daa82768-f08f-4775-80ca-734166101729"
	DomainTree_Group2Leaf1Name = api.DomainName("leaf1.group2.tree.testing.fixtures.rxp")
	DomainTree_Group2Leaf1     = domain.New(
		domain.WithUUID(DomainTree_Group2Leaf1UUID),
		domain.WithName(DomainTree_Group2Leaf1Name),
		domain.WithParent(DomainTree_Group2),
	)
	DomainTree_Group2Leaf2UUID = "a6846a47-590d-455f-b359-a4a8a0b3d236"
	DomainTree_Group2Leaf2Name = api.DomainName("leaf2.group2.tree.testing.fixtures.rxp")
	DomainTree_Group2Leaf2     = domain.New(
		domain.WithUUID(DomainTree_Group2Leaf2UUID),
		domain.WithName(DomainTree_Group2Leaf2Name),
		domain.WithParent(DomainTree_Group2),
	)
)
