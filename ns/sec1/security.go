package sec1

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

type Security struct {
	NameSpace jsonld.NameSpace `jsonld:"https://w3id.org/security/v1"`

	PublicKey opt.Optional[PublicKey] `json:"publicKey,omitempty"`
}
