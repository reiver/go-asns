package toot

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

//"toot": "http://joinmastodon.org/ns#"
type Toot struct {
	NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
	Prefix    jsonld.Prefix    `jsonld:"toot"`


	AttributionDomains opt.Optional[string] `jsonld:"attributionDomains,omitempty"`
	BlurHash           opt.Optional[string] `jsonld:"blurhash,omitempty"`
	Discoverable     opt.Optional[bool]   `jsonld:"discoverable,omitempty"`
	Emoji            opt.Optional[string] `jsonld:"Emoji,omitempty"`
	Featured         opt.Optional[bool]   `jsonld:"featured,omitempty"`
	FeaturedTags     []string             `jsonld:"featuredTags,omitempty"`
	FocalPoint       []any                `jsonld:"focalPoint,omitempty"`
	IdentityProof    opt.Optional[string]   `jsonld:"IdentityProof,omitempty"`
	Indexable        opt.Optional[bool]   `jsonld:"indexable,omitempty"`
	Memorial         opt.Optional[bool]   `jsonld:"memorial,omitempty"`
	VotersCount        opt.Optional[string] `jsonld:"votersCount,omitempty"`
	Suspended        opt.Optional[bool]   `jsonld:"suspended,omitempty"`
}
