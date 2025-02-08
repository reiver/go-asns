package act

import (
	"github.com/reiver/go-opt"
)

type Generator struct {
	Name opt.Optional[string] `json:"name,omitempty"`
	Type opt.Optional[string] `json:"type,omitempty"`
	URL  opt.Optional[string] `json:"url,omitempty"`
}
