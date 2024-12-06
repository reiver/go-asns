package as

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-jsonld"
)

// Object represents the ActivityStreams object name-space used in a JSON-LD document.
//
// Reference:
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-object
//
// Example usage:
//
//	import (
//		"github.com/reiver/go-act/ns/as"
//		"github.com/reiver/go-jsonld"
//	)
//	
//	// ...
//	
//	var object as.Object
//	
//	// ...
//	
//	bytes, err := jsonld.Marshal(object)
//
// More likely you would mix this with other JSON-LD name-spaces, with code similar to the following:
//
//	import (
//		"github.com/reiver/go-act/ns/ap"
//		"github.com/reiver/go-act/ns/as"
//		"github.com/reiver/go-act/ns/toot"
//		"github.com/reiver/go-jsonld"
//	)
//	
//	// ...
//	
//	var actor ap.Actor
//	var object as.Object
//	var toot toot.Toot
//	
//	// ...
//	
//	bytes, err := jsonld.Marshal(activitypub, object, toot)
type Object struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	Icon	opt.Optional[Icon]   `jsonld:"icon,omitempty"`
	Name    opt.Optional[string] `jsonld:"name,omitempty"`
	Summary opt.Optional[string] `jsonld:"summary,omitempty"`
	URL     opt.Optional[string] `jsonld:"url,omitempty"`
}
