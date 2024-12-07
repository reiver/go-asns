package toot

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

//"toot": "http://joinmastodon.org/ns#"
type Toot struct {
	NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
	Prefix    jsonld.Prefix    `jsonld:"toot"`

	AttributionDomains opt.Optional[string] `json:"attributionDomains,omitempty"`
	BlurHash           opt.Optional[string] `json:"blurhash,omitempty"`
	Discoverable       opt.Optional[bool]   `json:"discoverable,omitempty"`
	Emoji              opt.Optional[string] `json:"Emoji,omitempty"`
	Featured           opt.Optional[bool]   `json:"featured,omitempty"`
	FeaturedTags     []string               `json:"featuredTags,omitempty"`
	FocalPoint       []any                  `json:"focalPoint,omitempty"`
	IdentityProof      opt.Optional[string] `json:"IdentityProof,omitempty"`
	Indexable          opt.Optional[bool]   `json:"indexable,omitempty"`
	Memorial           opt.Optional[bool]   `json:"memorial,omitempty"`
	VotersCount        opt.Optional[string] `json:"votersCount,omitempty"`
	Suspended          opt.Optional[bool]   `json:"suspended,omitempty"`
}
