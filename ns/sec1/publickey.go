package sec1

import (
	"github.com/reiver/go-opt"
)

type PublicKey struct {
	ID           opt.Optional[string] `json:"id,omitempty"`
	Owner        opt.Optional[string] `json:"owner,omitempty"`
	PublicKeyPem opt.Optional[string] `json:"publicKeyPem,omitempty"`
}
