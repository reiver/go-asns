package act

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-orderedcollection
type OrderedCollection struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	Current     opt.Optional[string] `json:"current,omitempty"`                   // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-current
	First       opt.Optional[string] `json:"first,omitempty"`                     // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-first
	Icon        Icon                 `json:"icon,omitempty"`                      // https://www.w3.org/ns/activitystreams#icon
	ID          opt.Optional[string] `json:"id,omitempty"`                        // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-id
	Image       Image                `json:"image,omitempty"`                     // https://www.w3.org/ns/activitystreams#image
	Items       []any                `json:"items,omitempty"`                     // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-items
	Last        opt.Optional[string] `json:"last,omitempty"`                      // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-last
	Name        opt.Optional[string] `json:"name,omitempty"`                      // https://www.w3.org/ns/activitystreams#name
	Summary     opt.Optional[string] `json:"summary,omitempty"`                   // https://www.w3.org/ns/activitystreams#summary
	TotalItems  opt.Optional[uint64] `json:"totalItems,omitempty"`                // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-totalitems
	Type          json.Const[string] `json:"type" json.value:"OrderedCollection"` // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-type
}
