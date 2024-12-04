package toot

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

//"toot": "http://joinmastodon.org/ns#"
type Toot struct {
	NameSpace jsonld.NameSpace `jsonld:"http://joinmastodon.org/ns#"`
	Prefix    jsonld.Prefix    `jsonld:"toot"`

	Curve25519Key    opt.Optional[string] `jsonld:"Curve25519Key,omitempty"`
	CipherText       opt.Optional[string] `jsonld:"cipherText,omitempty"`
	Claim            opt.Optional[string] `jsonld:"claim,omitempty"`
	Device           opt.Optional[string] `jsonld:"Device,omitempty"`
	DeviceID         opt.Optional[string] `jsonld:"deviceId,omitempty"`
	Discoverable     opt.Optional[bool]   `jsonld:"discoverable,omitempty"`
	Ed25519Key       opt.Optional[string] `jsonld:"Ed25519Key,omitempty"`
	Ed25519Signature opt.Optional[string] `jsonld:"Ed25519Signature,omitempty"`
	Emoji            opt.Optional[string] `jsonld:"Emoji,omitempty"`
	EncryptedMessage opt.Optional[string] `jsonld:"EncryptedMessage,omitempty"`
	Featured         opt.Optional[bool]   `jsonld:"featured,omitempty"`
	FeaturedTags     []string             `jsonld:"featuredTags,omitempty"`
	FingerPrintKey   opt.Optional[string] `jsonld:"fingerprintKey,omitempty"`
	FocalPoint       []any                `jsonld:"focalPoint,omitempty"`
	IdentityKey      opt.Optional[string] `jsonld:"identityKey,omitempty"`
	Indexable        opt.Optional[bool]   `jsonld:"indexable,omitempty"`
	Memorial         opt.Optional[bool]   `jsonld:"memorial,omitempty"`
	MessageFranking  opt.Optional[string] `jsonld:"messageFranking,omitempty"`
	MessageType      opt.Optional[string] `jsonld:"messageType,omitempty"`
	PublicKeyBase64  opt.Optional[string] `jsonld:"publicKeyBase64,omitempty"`
	Suspended        opt.Optional[bool]   `jsonld:"suspended,omitempty"`
}
