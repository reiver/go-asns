package asns

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
)

// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-orderedcollectionpage
type OrderedCollectionPage struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	Icon         Icon                 `json:"icon,omitempty"`                          // https://www.w3.org/ns/activitystreams#icon
	ID           opt.Optional[string] `json:"id,omitempty"`                            // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-id
	Image        Image                `json:"image,omitempty"`                         // https://www.w3.org/ns/activitystreams#image
	Name         opt.Optional[string] `json:"name,omitempty"`                          // https://www.w3.org/ns/activitystreams#name
	Next         opt.Optional[string] `json:"next,omitempty"`                          // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-next
	OrderedItems []any                `json:"orderedItems"`                            // https://www.w3.org/TR/activitystreams-core/#h-paging
	PartOf       opt.Optional[string] `json:"partOf,omitempty"`                        // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-partof
	Prev         opt.Optional[string] `json:"prev,omitempty"`                          // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-prev
	StartIndex   opt.Optional[uint64] `json:"startIndex,omitempty"`                    // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-startindex
	Summary      opt.Optional[string] `json:"summary,omitempty"`                       // https://www.w3.org/ns/activitystreams#summary
	Type           json.Const[string] `json:"type" json.value:"OrderedCollectionPage"` // https://www.w3.org/TR/activitystreams-vocabulary/#dfn-type
}
