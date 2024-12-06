package as

import (
	"github.com/reiver/go-opt"
)

type Icon struct {
	Height opt.Optional[uint64] `jsonld:"height,omitempty"`
	Name   opt.Optional[string] `jsonld:"name,omitempty"`
	Type   opt.Optional[string] `jsonld:"type,omitempty"`
	URL    opt.Optional[string] `jsonld:"url,omitempty"`
	Width  opt.Optional[uint64] `jsonld:"width,omitempty"`
}
