package as

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// https://www.w3.org/ns/activitystreams#Person
type Person struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	AlsoKnownAs []string             `json:"alsoKnownAs,omitempty"`
	Attachment  []Attachment         `json:"attachment,omitempty"`     // https://www.w3.org/ns/activitystreams#attachment
	ID          opt.Optional[string] `json:"id,omitempty"`
	Image	    Image                `json:"image,omitempty"`          // https://www.w3.org/ns/activitystreams#image
	Icon	    Icon                 `json:"icon,omitempty"`           // https://www.w3.org/ns/activitystreams#icon
	MediaType   opt.Optional[string] `json:"mediaType,omitempty"`      // https://www.w3.org/ns/activitystreams#mediaType
	MovedTo     opt.Optional[string] `json:"movedTo,omitempty"`
	Name        opt.Optional[string] `json:"name,omitempty"`           // https://www.w3.org/ns/activitystreams#name
	Summary     opt.Optional[string] `json:"summary,omitempty"`        // https://www.w3.org/ns/activitystreams#summary
	Tag         []HashTag            `json:"tag,omitempty"`            // https://www.w3.org/ns/activitystreams#tag
	Type          json.Const[string] `json:"type" json.value:"Person"`
	URL         opt.Optional[string] `json:"url,omitempty"`            // https://www.w3.org/ns/activitystreams#url
}
