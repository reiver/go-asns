package as_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-act/ns/as"
)

func TestIcon_MarshalJSONLD(t *testing.T) {

	tests := []struct{
		Icon as.Icon
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Icon: as.Icon{
				Height: opt.Something(uint64(123)),
			},
			Expected: []byte(`{"height":123}`),
		},
		{
			Icon: as.Icon{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"name":"apple banana cherry"}`),
		},
		{
			Icon: as.Icon{
				Type: opt.Something("image/png"),
			},
			Expected: []byte(`{"type":"image/png"}`),
		},
		{
			Icon: as.Icon{
				URL: opt.Something("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="),
			},
			Expected: []byte(`{"url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="}`),
		},
		{
			Icon: as.Icon{
				Width: opt.Something(uint64(4567)),
			},
			Expected: []byte(`{"width":4567}`),
		},



		{
			Icon: as.Icon{
				Height: opt.Something(uint64(123)),
				Name: opt.Something("apple banana cherry"),
				Type: opt.Something("image/png"),
				URL: opt.Something("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="),
				Width: opt.Something(uint64(4567)),
			},
			Expected: []byte(`{"height":123,"name":"apple banana cherry","type":"image/png","url":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==","width":4567}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.Marshal(test.Icon)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ICON: %#v", test.Icon)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual jsonld-marshaled %T is not what was expected.", testNumber, test.Icon)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("ICON: %#v", test.Icon)
			continue
		}
	}
}
