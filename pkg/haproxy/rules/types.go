package rules

// Order matters !
// Rules will be evaluated by HAProxy in the defined order.
type Type int

//nolint: golint,stylecheck
const (
	REQ_ACCEPT_CONTENT Type = iota
	REQ_INSPECT_DELAY
	REQ_PROXY_PROTOCOL
	REQ_SET_VAR
	REQ_SET_SRC
	REQ_DENY
	REQ_TRACK
	REQ_AUTH
	REQ_RATELIMIT
	REQ_CAPTURE
	REQ_REDIRECT
	REQ_FORWARDED_PROTO
	REQ_SET_HEADER
	REQ_SET_HOST
	REQ_PATH_REWRITE
	RES_SET_HEADER
)

var constLookup = map[Type]string{
	REQ_ACCEPT_CONTENT:  "REQ_ACCEPT_CONTENT",
	REQ_INSPECT_DELAY:   "REQ_INSPECT_DELAY",
	REQ_PROXY_PROTOCOL:  "REQ_PROXY_PROTOCOL",
	REQ_SET_VAR:         "REQ_SET_VAR",
	REQ_SET_SRC:         "REQ_SET_SRC",
	REQ_DENY:            "REQ_DENY",
	REQ_TRACK:           "REQ_TRACK",
	REQ_AUTH:            "REQ_AUTH",
	REQ_RATELIMIT:       "REQ_RATELIMIT",
	REQ_CAPTURE:         "REQ_CAPTURE",
	REQ_REDIRECT:        "REQ_REDIRECT",
	REQ_FORWARDED_PROTO: "REQ_FORWARDED_PROTO",
	REQ_SET_HEADER:      "REQ_SET_HEADER",
	REQ_SET_HOST:        "REQ_SET_HOST",
	REQ_PATH_REWRITE:    "REQ_PATH_REWRITE",
	RES_SET_HEADER:      "RES_SET_HEADER",
}