package as

import (
	"github.com/reiver/go-opt"
)

type Icon struct {
	Height opt.Optional[uint64] `json:"height,omitempty"`
	Name   opt.Optional[string] `json:"name,omitempty"`
	Type   opt.Optional[string] `json:"type,omitempty"`
	URL    opt.Optional[string] `json:"url,omitempty"`
	Width  opt.Optional[uint64] `json:"width,omitempty"`
}
