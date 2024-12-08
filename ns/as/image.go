package as

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-opt"
)

type Image struct {
	Audiences            []Audience  `json:"audience,omitempty"`
	Generator              Generator `json:"generator,omitempty"`
	Height    opt.Optional[uint64]   `json:"height,omitempty"`
	MediaType opt.Optional[string]   `json:"mediaType,omitempty"`
	Location               Location  `json:"location,omitempty"`
	Name      opt.Optional[string]   `json:"name,omitempty"`
	Published opt.Optional[string]   `json:"published,omitempty"`
	Type        json.Const[string]   `json:"type" json.value:"Image"`
	Updated   opt.Optional[string]   `json:"updated,omitempty"`
	URL       opt.Optional[string]   `json:"url,omitempty"`
	Width     opt.Optional[uint64]   `json:"width,omitempty"`
}
