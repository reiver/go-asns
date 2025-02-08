package act_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-act"
)

func TestIcon_marshal(t *testing.T) {

	const context string =
		`"@context":{`+
			`"as":"https://www.w3.org/ns/activitystreams"`+
			`,`+
			`"audience":"as:audience"`+
			`,`+
			`"generator":"as:generator"`+
			`,`+
			`"height":"as:height"`+
			`,`+
			`"location":"as:location"`+
			`,`+
			`"mediaType":"as:mediaType"`+
			`,`+
			`"name":"as:name"`+
			`,`+
			`"published":"as:published"`+
			`,`+
			`"type":"as:type"`+
			`,`+
			`"updated":"as:updated"`+
			`,`+
			`"url":"as:url"`+
			`,`+
			`"width":"as:width"`+
		`}`

	type Value struct {
		Icon act.Icon `json:"icon,omitempty"`
	}

	tests := []struct{
		Value Value
		Expected []byte
	}{
		{
//			Expected: []byte(`{}`),
			Expected: []byte(`{`+context+`,"type":"Image"}`),
		},



		{
			Value: Value{
				Icon: act.Icon{
					Height: opt.Something(uint64(123)),
				},
			},
//			Expected: []byte(`{"icon":{"height":123,"type":"Image"}}`),
			Expected: []byte(`{`+context+`,"icon":{"height":123,"type":"Image"}}`),
		},
		{
			Value: Value{
				Icon: act.Icon{
					Name: opt.Something("apple banana cherry"),
				},
			},
//			Expected: []byte(`{"icon":{"name":"apple banana cherry","type":"Image"}}`),
			Expected: []byte(`{`+context+`,"icon":{"name":"apple banana cherry","type":"Image"}}`),
		},
		{
			Value: Value{
				Icon: act.Icon{
					MediaType: opt.Something("image/png"),
				},
			},
//			Expected: []byte(`{"icon":{"mediaType":"image/png","type":"Image"}}`),
			Expected: []byte(`{`+context+`,"icon":{"mediaType":"image/png","type":"Image"}}`),
		},
		{
			Value: Value{
				Icon: act.Icon{
					URL: opt.Something("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="),
				},
			},
//			Expected: []byte(`{"icon":{"type":"Image","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="}}`),
			Expected: []byte(`{`+context+`,"icon":{"type":"Image","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="}}`),
		},
		{
			Value: Value{
				Icon: act.Icon{
					Width: opt.Something(uint64(4567)),
				},
			},
//			Expected: []byte(`{"icon":{"type":"Image","width":4567}}`),
			Expected: []byte(`{`+context+`,"icon":{"type":"Image","width":4567}}`),
		},



		{
			Value: Value{
				Icon: act.Icon{
					Height: opt.Something(uint64(123)),
					MediaType: opt.Something("image/png"),
					Name: opt.Something("apple banana cherry"),
					URL: opt.Something("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="),
					Width: opt.Something(uint64(4567)),
				},
			},
//			Expected: []byte(`{"icon":{"height":123,"mediaType":"image/png","name":"apple banana cherry","type":"Image","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==","width":4567}}`),
			Expected: []byte(`{`+context+`,"icon":{"height":123,"mediaType":"image/png","name":"apple banana cherry","type":"Image","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==","width":4567}}`),
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
