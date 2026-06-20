package errors

var (
	ErrMissingSelector             = New("missing selector", WithWrap(ErrInvalidReadRequest))
	ErrInvalidSelector             = New("invalid selector", WithWrap(ErrInvalidReadRequest))
	ErrSelectorKindVersionRequired = New("kindversion required", WithWrap(ErrInvalidSelector))
	ErrSelectorKindRequired        = New("kind required", WithWrap(ErrInvalidSelector))
	ErrSelectorUUIDRequired        = New("uuid required", WithWrap(ErrInvalidSelector))
	ErrSelectorUUIDOrNameRequired  = New("uuid or name required", WithWrap(ErrInvalidSelector))
	ErrSelectorDomainRequired      = New("domain required", WithWrap(ErrInvalidSelector))
	ErrSelectorNameRequired        = New("name required", WithWrap(ErrInvalidSelector))
	ErrSelectorSystemMismatched    = New("selector system does not match domain system", WithWrap(ErrInvalidSelector))
)
