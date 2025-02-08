package as

import (
	"github.com/reiver/go-opt"
)

type Attachment struct {
	Type  opt.Optional[string] `json:"type,omitempty"`
	Name  opt.Optional[string] `json:"name,omitempty"`
	Value opt.Optional[string] `json:"value,omitempty"`
}
