package query

// Option controls the behaviour of [Query]
type Option func(*Options)

// Options controls how a call to Query behaves.
type Options struct {
	// limit is the max number of items to return. Note that the default (0)
	// does not mean return all items. The server will *always* bound the
	// number of returned items to some default value. Callers can check what
	// that default limit is by examining Result.Options.Limit()
	limit int
	// continueFrom contains the UUID of the last item on a previous "page" of
	// results returned from a call to Query. This value can be passed in
	// subsequent calls to Query to "continue" the list from that item.
	continueFrom string
}

// Limit returns the max number of items to return.
func (o Options) Limit() int {
	return o.limit
}

// ContinueFrom returns the UUID of the last item on a previous "page" of
// results returned from a call to Query. This value can be passed in
// subsequent calls to Query to "continue" the list from that item.
func (o Options) ContinueFrom() string {
	return o.continueFrom
}

// Limit specifies the max number of items to return. Note that the default
// (0) does not mean return all items. The server will *always* bound the
// number of returned items to some default value. Callers can check what that
// default limit is by examining Result.Options.Limit()
func Limit(limit int) Option {
	return func(o *Options) {
		o.limit = limit
	}
}

// ContinueFrom specifies the UUID of the last item on a previous "page" of
// results returned from a call to Query. Pass this value in subsequent calls
// to Query to "continue" the list from that item.
func ContinueFrom(continueFrom string) Option {
	return func(o *Options) {
		o.continueFrom = continueFrom
	}
}

// NewOptions returns a new Options given zero or more Options.
func NewOptions(opts ...Option) Options {
	o := Options{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
