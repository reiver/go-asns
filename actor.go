package act

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
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	EndPoints                  opt.Optional[string] `json:"endpoints,omitempty"`
	Following                  opt.Optional[string] `json:"following,omitempty"`
	Followers                  opt.Optional[string] `json:"followers,omitempty"`

	// inbox
	//
	// A reference to an ActivityStreams (1) OrderedCollection (2) comprised of all the messages received by the actor; see 5.2 Inbox (3).
	//
	// (1) https://www.w3.org/TR/activitystreams-core/
	//
	// (2) https://www.w3.org/TR/activitystreams-vocabulary/#dfn-orderedcollection
	//
	// (3) https://www.w3.org/TR/activitypub/#inbox
	Inbox                      opt.Optional[string] `json:"inbox,omitempty"`

	Liked                      opt.Optional[string] `json:"liked,omitempty"`
	Shares                     opt.Optional[string] `json:"shares,omitempty"`
	Likes                      opt.Optional[string] `json:"likes,omitempty"`
	OauthAuthorizationEndPoint opt.Optional[string] `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndPoint         opt.Optional[string] `json:"oauthTokenEndpoint,omitempty"`

	// outbox
	//
	// An ActivityStreams (1) OrderedCollection (2) comprised of all the messages produced by the actor; see 5.1 Outbox (3).
	//
	// (1) https://www.w3.org/TR/activitystreams-core/
	//
	// (2) https://www.w3.org/TR/activitystreams-vocabulary/#dfn-orderedcollection
	//
	// (3) https://www.w3.org/TR/activitypub/#outbox
	Outbox                     opt.Optional[string] `json:"outbox,omitempty"`

	// preferredUsername
	//
	// A short username which may be used to refer to the actor, with no uniqueness guarantees.
	PreferredUserName          opt.Optional[string] `json:"preferredUsername,omitempty"`

	ProvideClientKey           opt.Optional[string] `json:"provideClientKey,omitempty"`
	ProxyURL                   opt.Optional[string] `json:"proxyUrl,omitempty"`

	// sharedInbox
	//
	// An optional endpoint used for wide delivery of publicly addressed activities and activities sent to followers. (1)
	// sharedInbox endpoints SHOULD also be publicly readable OrderedCollection objects containing objects addressed to the Public (2) special collection. Reading from the sharedInbox endpoint MUST NOT present objects which are not addressed to the Public endpoint.
	//
	// (1) https://www.w3.org/TR/activitypub/#shared-inbox-delivery
	//
	// (2) https://www.w3.org/TR/activitypub/#public-addressing
	SharedInbox                opt.Optional[string] `json:"sharedInbox,omitempty"`

	SignClientKey              opt.Optional[string] `json:"signClientKey,omitempty"`
	Source                     opt.Optional[string] `json:"source,omitempty"`
	Streams                    opt.Optional[string] `json:"streams,omitempty"`
	UploadMedia                opt.Optional[string] `json:"uploadMedia,omitempty"`
}
