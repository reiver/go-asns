package ap

import (
	"github.com/reiver/go-opt"
)

// Actor represents the ActivityPub actor object name-space used in a JSON-LD document.
//
// Reference:
// https://www.w3.org/TR/activitypub/#actor-objects
//
// Example usage:
//
//	import (
//		"github.com/reiver/go-act/ns/ap"
//		"github.com/reiver/go-jsonld"
//	)
//	
//	// ...
//	
//	var actor ap.Actor
//	
//	// ...
//	
//	bytes, err := jsonld.Marshal(actor)
//
// More likely you would mix this with other JSON-LD name-spaces, with code similar to the following:
//
//	import (
//		"github.com/reiver/go-act/ns/ap"
//		"github.com/reiver/go-act/ns/toot"
//		"github.com/reiver/go-jsonld"
//	)
//	
//	// ...
//	
//	var actor ap.Actor
//	var toot toot.Toot
//	
//	// ...
//	
//	bytes, err := jsonld.Marshal(activitypub, toot)
type Actor struct {
	EndPoints                  opt.Optional[string] `jsonld:"endpoints,omitempty"`
	Following                  opt.Optional[string] `jsonld:"following,omitempty"`
	Followers                  opt.Optional[string] `jsonld:"followers,omitempty"`
	Inbox                      opt.Optional[string] `jsonld:"inbox,omitempty"`
	Liked                      opt.Optional[string] `jsonld:"liked,omitempty"`
	Shares                     opt.Optional[string] `jsonld:"shares,omitempty"`
	Likes                      opt.Optional[string] `jsonld:"likes,omitempty"`
	OauthAuthorizationEndPoint opt.Optional[string] `jsonld:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndPoint         opt.Optional[string] `jsonld:"oauthTokenEndpoint,omitempty"`
	Outbox                     opt.Optional[string] `jsonld:"outbox,omitempty"`
	PreferredUserName          opt.Optional[string] `jsonld:"preferredUsername,omitempty"`
	ProvideClientKey           opt.Optional[string] `jsonld:"provideClientKey,omitempty"`
	ProxyURL                   opt.Optional[string] `jsonld:"proxyUrl,omitempty"`
	SharedInbox                opt.Optional[string] `jsonld:"sharedInbox,omitempty"`
	SignClientKey              opt.Optional[string] `jsonld:"signClientKey,omitempty"`
	Source                     opt.Optional[string] `jsonld:"source,omitempty"`
	Streams                    opt.Optional[string] `jsonld:"streams,omitempty"`
	UploadMedia                opt.Optional[string] `jsonld:"uploadMedia,omitempty"`
}
