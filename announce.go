package asns

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

type Announce struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	Actor       jsonld.Strings       `json:"actor,omitempty"`
	AlsoKnownAs []string             `json:"alsoKnownAs,omitempty"`
	Attachment  []Attachment         `json:"attachment,omitempty"`
	Content     opt.Optional[string] `json:"content,omitempty"`
	ID          opt.Optional[string] `json:"id,omitempty"`
	Image	    Image                `json:"image,omitempty"`
	Icon	    Icon                 `json:"icon,omitempty"`
	MediaType   opt.Optional[string] `json:"mediaType,omitempty"`
	MovedTo     opt.Optional[string] `json:"movedTo,omitempty"`
	Name        opt.Optional[string] `json:"name,omitempty"`
	Object      opt.Optional[string] `json:"object,omitempty"`
	Published   opt.Optional[string] `json:"published,omitempty"`
	Summary     opt.Optional[string] `json:"summary,omitempty"`
	Tag         []HashTag            `json:"tag,omitempty"`
	Target      opt.Optional[string] `json:"target,omitempty"`
	To          []string             `json:"to,omitempty"`
	Type        json.Const[string]   `json:"type" json.value:"Announce"`
	URL         opt.Optional[string] `json:"url,omitempty"`
}
