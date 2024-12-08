package as

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

type Link struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	Height    opt.Optional[string] `json:"height,omitempty,bare"`
	HRef      opt.Optional[string] `json:"href,omitempty"`
	HRefLang  opt.Optional[string] `json:"hreflang,omitempty"`
	MediaType opt.Optional[string] `json:"mediaType,omitempty"`
	Name      opt.Optional[string] `json:"name,omitempty"`
	Rel       opt.Optional[string] `json:"rel,omitempty"`
	Type        json.Const[string] `json:"type,omitempty" json.value:"Link"`
	Width     opt.Optional[string] `json:"width,omitempty,bare"`
}

func (receiver Link) IsEmpty() bool {
	var nada Link
	return nada == receiver
}
