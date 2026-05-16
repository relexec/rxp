package option

// Option can be used to control how [ObjectLister.ObjectList] behaves.
type Option func(*Options)

// Options controls how a call to [ObjectLister.ObjectList] behaves.
type Options struct {
	// limit is the max number of Objects to return. Note that the default (0)
	// does not mean return all Objects. The server will *always* bound the
	// number of returned Objects to some default value. Callers can check what
	// that default limit is by examining ObjectListResults.Options.Limit()
	limit int
	// marker contains the UUID of the last Object on a previous "page" of
	// results returned from a call to ObjectList. This value can be passed in
	// subsequent calls to ObjectList to "continue" the list from that Object.
	marker string
}

// Limit returns the max number of Objects to return.
func (o Options) Limit() int {
	return o.limit
}

// Marker returns the Marker of the Object's Spec to read.
func (o Options) Marker() string {
	return o.marker
}

// WithLimit specifies the max number of Objects to return. Note that the
// default (0) does not mean return all Objects. The server will *always* bound
// the number of returned Objects to some default value. Callers can check what
// that default limit is by examining ObjectListResults.Options.Limit()
func WithLimit(limit int) Option {
	return func(o *Options) {
		o.limit = limit
	}
}

// WithMarker specifies the UUID of the last Object on a previous "page" of
// results returne from a call to ObjectList. Pass this value in subsequent
// calls to ObjectList to "continue" the list from that Object.
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
