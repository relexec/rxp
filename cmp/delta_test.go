package cmp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
)

type Obj struct {
	Spec   ObjSpec
	Status ObjStatus
}

type ObjSpec struct {
	Replicas   int
	Generation int
}

type ObjStatus struct {
	ReadyReplicas int
}

var (
	objA = Obj{
		Spec: ObjSpec{
			Replicas:   1,
			Generation: 1,
		},
		Status: ObjStatus{
			ReadyReplicas: 0,
		},
	}
	// B is different from A in spec.replicas and status.readyReplicas
	objB = Obj{
		Spec: ObjSpec{
			Replicas:   2,
			Generation: 1,
		},
		Status: ObjStatus{
			ReadyReplicas: 1,
		},
	}
	// C is different from A only in status.readyReplicas
	// C is different from B in spec.replicas and status.readyReplicas
	objC = Obj{
		Spec: ObjSpec{
			Replicas:   1,
			Generation: 1,
		},
		Status: ObjStatus{
			ReadyReplicas: 0,
		},
	}
)

var (
	fpSpecReplicas        = fieldpath.FieldPath{"spec", "replicas"}
	fpSpecGeneration      = fieldpath.FieldPath{"spec", "generation"}
	fpStatusReadyReplicas = fieldpath.FieldPath{"status", "readyReplicas"}
)

var (
	diffA_FromZero_SpecReplicas = cmp.NewDifference(
		fpSpecReplicas,
		cmp.DifferenceTypeAdd,
		nil,
		1,
	)
	diffA_FromZero_SpecGeneration = cmp.NewDifference(
		fpSpecGeneration,
		cmp.DifferenceTypeAdd,
		nil,
		1,
	)
	diffA_FromZero_StatusReadyReplicas = cmp.NewDifference(
		fpStatusReadyReplicas,
		cmp.DifferenceTypeAdd,
		nil,
		1,
	)
	diffAB_SpecReplicas = cmp.NewDifference(
		fpSpecReplicas,
		cmp.DifferenceTypeModify,
		1,
		2,
	)
	diffAB_StatusReadyReplicas = cmp.NewDifference(
		fpStatusReadyReplicas,
		cmp.DifferenceTypeModify,
		0,
		1,
	)
	diffAC_StatusReadyReplicas = cmp.NewDifference(
		fpStatusReadyReplicas,
		cmp.DifferenceTypeModify,
		0,
		1,
	)
	diffBC_SpecReplicas = cmp.NewDifference(
		fpSpecReplicas,
		cmp.DifferenceTypeModify,
		2,
		1,
	)
	diffBC_StatusReadyReplicas = cmp.NewDifference(
		fpStatusReadyReplicas,
		cmp.DifferenceTypeModify,
		1,
		0,
	)
)

var (
	emptyDelta = cmp.NewDelta()
	deltaAB    = cmp.NewDelta(
		diffAB_SpecReplicas,
		diffAB_StatusReadyReplicas,
	)
	deltaAC = cmp.NewDelta(
		diffAC_StatusReadyReplicas,
	)
	deltaBC = cmp.NewDelta(
		diffBC_SpecReplicas,
		diffBC_StatusReadyReplicas,
	)
	deltaA_FromZero = cmp.NewDelta(
		diffA_FromZero_SpecReplicas,
		diffA_FromZero_SpecGeneration,
		diffA_FromZero_StatusReadyReplicas,
	)
)

func TestDelta_Different(t *testing.T) {
	cases := []struct {
		name    string
		subject cmp.Delta
		exp     bool
	}{
		{
			"empty delta",
			emptyDelta,
			false,
		},
		{
			"non-empty delta",
			deltaAB,
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(c.exp, c.subject.Different())
		})
	}
}

func TestDelta_DifferentAt(t *testing.T) {
	cases := []struct {
		name    string
		subject cmp.Delta
		path    fieldpath.FieldPath
		exp     bool
	}{
		{
			"empty delta",
			emptyDelta,
			fpSpecReplicas,
			false,
		},
		{
			"AB spec.replicas",
			deltaAB,
			fpSpecReplicas,
			true,
		},
		{
			"AB spec.generation",
			deltaAB,
			fpSpecGeneration,
			false,
		},
		{
			"AB status.readyReplicas",
			deltaAB,
			fpStatusReadyReplicas,
			true,
		},
		{
			"AC spec.replicas",
			deltaAC,
			fpSpecReplicas,
			false,
		},
		{
			"AC spec.generation",
			deltaAC,
			fpSpecGeneration,
			false,
		},
		{
			"AC status.readyReplicas",
			deltaAC,
			fpStatusReadyReplicas,
			true,
		},
		{
			"BC spec.replicas",
			deltaBC,
			fpSpecReplicas,
			true,
		},
		{
			"BC spec.generation",
			deltaBC,
			fpSpecGeneration,
			false,
		},
		{
			"BC status.readyReplicas",
			deltaBC,
			fpStatusReadyReplicas,
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(c.exp, c.subject.DifferentAt(c.path))
		})
	}
}

func TestDelta_DifferentExcept(t *testing.T) {
	cases := []struct {
		name    string
		subject cmp.Delta
		paths   []fieldpath.FieldPath
		exp     bool
	}{
		{
			"empty delta",
			emptyDelta,
			[]fieldpath.FieldPath{fpSpecReplicas},
			false,
		},
		{
			"AB except spec.replicas",
			deltaAB,
			[]fieldpath.FieldPath{fpSpecReplicas},
			true, // because also diff at status.readyReplicas
		},
		{
			"AB except spec.replicas and status.readyReplicas",
			deltaAB,
			[]fieldpath.FieldPath{fpSpecReplicas, fpStatusReadyReplicas},
			false, // because above paths are the only diffs
		},
		{
			"AC except status.readyReplicas",
			deltaAC,
			[]fieldpath.FieldPath{fpStatusReadyReplicas},
			false, // because above path is the only diffs
		},
		{
			"no supplied field paths and non-empty delta",
			deltaAB,
			nil,
			true, // because there are differences
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(c.exp, c.subject.DifferentExcept(c.paths...))
		})
	}
}

func TestDelta_Marshal(t *testing.T) {
	cases := []struct {
		name    string
		subject cmp.Delta
		exp     string
		expErr  string
	}{
		{
			"empty delta",
			emptyDelta,
			"",
			"",
		},
		{
			"from zero",
			deltaA_FromZero,
			"[{\"path\":\"spec.replicas\",\"type\":0,\"to\":1},{\"path\":\"spec.generation\",\"type\":0,\"to\":1},{\"path\":\"status.readyReplicas\",\"type\":0,\"to\":1}]",
			"",
		},
		{
			"delta changes A to B",
			deltaAB,
			"[{\"path\":\"spec.replicas\",\"type\":2,\"from\":1,\"to\":2},{\"path\":\"status.readyReplicas\",\"type\":2,\"from\":0,\"to\":1}]",
			"",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			got, err := c.subject.MarshalJSON()
			if c.expErr != "" {
				require.ErrorContains(err, c.expErr)
			} else {
				require.Nil(err)
				require.Equal(c.exp, string(got))
			}
		})
	}
}
