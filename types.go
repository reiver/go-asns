package asns

import (
	gobytes "bytes"
	gojson "encoding/json"
	"strings"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-json"
	"github.com/reiver/go-opt"
)

var null []byte = []byte{'n','u','l','l'}

const typesSeparator string = "\x1F"

// Types represents an ActivityStreams "type" field.
type Types struct {
	optional opt.Optional[string]
}

func NoTypes() Types {
	return Types{}
}

func SomeType(value string) Types {
	return Types{
		optional: opt.Something(value),
	}
}

func SomeTypes(values ...string) Types {
	var str string = strings.Join(values, typesSeparator)

	return Types{
		optional: opt.Something(str),
	}
}

func (receiver Types) Types() []string {
	value, found := receiver.optional.Get()
	if !found {
		return []string{}
	}

	return strings.Split(value, typesSeparator)
}

func (receiver Types) MarshalJSON() ([]byte, error) {
	types := receiver.Types()

	switch len(types) {
	case 0:
		return append([]byte(nil), null...), nil
	case 1:
		return json.Marshal(types[0])
	default:
		return json.Marshal(types)
	}
}

func (receiver *Types) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if len(bytes) <= 0 {
		return errEmptyBytes
	}

	var byte0 byte = bytes[0]

	switch {
	case gobytes.Equal(null, bytes):
		*receiver = NoTypes()
		return nil
	case '"' == byte0:
		var str string

		err := gojson.Unmarshal(bytes, &str)
		if nil != err {
			return erorr.Errorf("asns: problem unmarshaling-json into string: %w", err)
		}

		*receiver = SomeType(str)
		return nil
	case '[' == byte0:
		var arr []string = []string{}

		err := gojson.Unmarshal(bytes, &arr)
		if nil != err {
			return erorr.Errorf("asns: problem unmarshaling-json into []string: %w", err)
		}

		*receiver = SomeTypes(arr...)
		return nil
	default:
		return erorr.Errorf("asns: cannot unmarshal-json something that starts with %q (%X)", byte0, byte0)
	}
}
