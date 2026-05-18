package option

// Option can be used to control how List behaves.
type Option func(*Options)

// Options controls how a call to List behaves.
type Options struct {
	// limit is the max number of items to return. Note that the default (0)
	// does not mean return all items. The server will *always* bound the
	// number of returned items to some default value. Callers can check what
	// that default limit is by examining Result.Options.Limit()
	limit int
	// marker contains the UUID of the last item on a previous "page" of
	// results returned from a call to List. This value can be passed in
	// subsequent calls to List to "continue" the list from that item.
	marker string
}

// Limit returns the max number of items to return.
func (o Options) Limit() int {
	return o.limit
}

// Marker returns the UUID of the last item on a previous "page" of results
// returned from a call to List. This value can be passed in subsequent calls
// to List to "continue" the list from that item.
func (o Options) Marker() string {
	return o.marker
}

// WithLimit specifies the max number of items to return. Note that the default
// (0) does not mean return all items. The server will *always* bound the
// number of returned items to some default value. Callers can check what that
// default limit is by examining Result.Options.Limit()
func WithLimit(limit int) Option {
	return func(o *Options) {
		o.limit = limit
	}
}

// WithMarker specifies the UUID of the last item on a previous "page" of
// results returne from a call to List. Pass this value in subsequent calls to
// List to "continue" the list from that item.
func WithMarker(marker string) Option {
	return func(o *Options) {
		o.marker = marker
	}
}

// New returns a new Options given zero or more Option modifiers.
func New(opts ...Option) Options {
	o := Options{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
