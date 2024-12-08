package as

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-opt"
)

type HashTag struct {
	HRef opt.Optional[string] `json:"href,omitempty"`
	Name opt.Optional[string] `json:"name,omitempty"`
	Type   json.Const[string] `json:"type,omitempty" json.value:"Hashtag"`
}
