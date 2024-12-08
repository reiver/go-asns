package as_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-act"
	"github.com/reiver/go-act/ns/as"
)

func TestIcon_marshal(t *testing.T) {

	type Value struct {
		Icon as.Icon `json:"icon,omitempty"`
	}

	tests := []struct{
		Value Value
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Value: Value{
				Icon: as.Icon{
					Height: opt.Something(uint64(123)),
				},
			},
			Expected: []byte(`{"icon":{"height":123,"type":"Image"}}`),
		},
		{
			Value: Value{
				Icon: as.Icon{
					Name: opt.Something("apple banana cherry"),
				},
			},
			Expected: []byte(`{"icon":{"name":"apple banana cherry","type":"Image"}}`),
		},
		{
			Value: Value{
				Icon: as.Icon{
					MediaType: opt.Something("image/png"),
				},
			},
			Expected: []byte(`{"icon":{"mediaType":"image/png","type":"Image"}}`),
		},
		{
			Value: Value{
				Icon: as.Icon{
					URL: opt.Something("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="),
				},
			},
			Expected: []byte(`{"icon":{"type":"Image","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="}}`),
		},
		{
			Value: Value{
				Icon: as.Icon{
					Width: opt.Something(uint64(4567)),
				},
			},
			Expected: []byte(`{"icon":{"type":"Image","width":4567}}`),
		},



		{
			Value: Value{
				Icon: as.Icon{
					Height: opt.Something(uint64(123)),
					MediaType: opt.Something("image/png"),
					Name: opt.Something("apple banana cherry"),
					URL: opt.Something("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="),
					Width: opt.Something(uint64(4567)),
				},
			},
			Expected: []byte(`{"icon":{"height":123,"mediaType":"image/png","name":"apple banana cherry","type":"Image","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==","width":4567}}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := act.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
