package asns

import (
	"github.com/reiver/go-jsonld"
)

func Marshal(values ...any) ([]byte, error) {
	return jsonld.Marshal(values...)
}

