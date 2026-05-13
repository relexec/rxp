package metrics

import (
	"go.opentelemetry.io/otel/attribute"

	"github.com/relexec/rxp/types"
)

const (
	AttributeNameErrCode = "error.code"
)

type hasStatusCode interface {
	StatusCode() int
}

// AttributeErrCode returns the error code attribute KeyValue with the value of
// the supplied error's code.
func AttributeErrCode(err error) attribute.KeyValue {
	code := 500
	if hc, ok := err.(hasStatusCode); ok {
		code = hc.StatusCode()
	}
	return attribute.Int(AttributeNameErrCode, code)
}

type TargetType string

const (
	TargetTypeSystem    TargetType = "system"
	TargetTypeKind      TargetType = "kind"
	TargetTypeDomain    TargetType = "domain"
	TargetTypeNamespace TargetType = "namespace"
	TargetTypeMeta      TargetType = "meta"
	TargetTypeObject    TargetType = "object"
)

const (
	AttributeNameTargetType = "target.type"
)

// AttributeTargetType returns the target type attribute KeyValue with the
// value of the supplied target type.
func AttributeTargetType(tt TargetType) attribute.KeyValue {
	return attribute.String(AttributeNameTargetType, string(tt))
}

const (
	AttributeNameKindVersion = "kindversion"
)

// AttributeKindVersion returns the kindversion attribute KeyValue with the
// value of the supplied KindVersion.
func AttributeKindVersion(kv types.KindVersion) attribute.KeyValue {
	return attribute.String(AttributeNameKindVersion, string(kv))
}
