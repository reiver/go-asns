package as

import (
	"github.com/reiver/go-opt"
)

type Location struct {
	Altitude  opt.Optional[string] `json:"altitude,omitempty,bare"`
	Latitude  opt.Optional[string] `json:"latitude,omitempty,bare"`
	Longitude opt.Optional[string] `json:"longitude,omitempty,bare"`
	Name      opt.Optional[string] `json:"name,omitempty"`
	Type      opt.Optional[string] `json:"type,omitempty"`
	Units     opt.Optional[string] `json:"units,omitempty"`
	URL       opt.Optional[string] `json:"url,omitempty"`
}
