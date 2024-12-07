package as_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-act/ns/as"
)

func TestObject_MarshalJSONLD(t *testing.T) {

	const context string =
		`"@context":{`+
			`"as":"https://www.w3.org/ns/activitystreams"`+
			`,`+
			`"icon":"as:icon"`+
			`,`+
			`"name":"as:name"`+
			`,`+
			`"summary":"as:summary"`+
			`,`+
			`"url":"as:url"`+
		`}`

	tests := []struct{
		Value as.Object
		Expected []byte
	}{
		// 0
		{
			Value: as.Object{},
			Expected: []byte(`{`+context+`}`),
		},



		// 1
		{
			Value: as.Object{
				Icon: opt.Something(as.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				}),
			},
			Expected: []byte(`{`+context+`,"icon":{"url":"http://example.com/img/icon.png"}}`),
		},
		// 2
		{
			Value: as.Object{
				Icon: opt.Something(as.Icon{
					Height: opt.Something(uint64(123)),
					Name: opt.Something("apple banana cherry"),
					Type: opt.Something("image/png"),
					URL: opt.Something("http://example.com/img/icon.png"),
					Width: opt.Something(uint64(456)),
				}),
			},
			Expected: []byte(`{`+context+`,"icon":{"height":123,"name":"apple banana cherry","type":"image/png","url":"http://example.com/img/icon.png","width":456}}`),
		},
		// 3
		{
			Value: as.Object{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"name":"apple banana cherry"}`),
		},
		// 4
		{
			Value: as.Object{
				Summary: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"summary":"apple banana cherry"}`),
		},
		// 5
		{
			Value: as.Object{
				URL: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"url":"apple banana cherry"}`),
		},



		// 6
		{
			Value: as.Object{
				Icon: opt.Something(as.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				}),
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"icon":{`+
						`"url":"http://example.com/img/icon.png"`+
					`}`+
					`,`+
					`"name":"apple banana cherry"`+
				`}`),
		},



		// 7
		{
			Value: as.Object{
				Name:    opt.Something("apple"),
				Summary: opt.Something("banana"),
				URL:     opt.Something("cherry"),
			},
			Expected: []byte(`{`+context+`,"name":"apple","summary":"banana","url":"cherry"}`),
		},
		// 8
		{
			Value: as.Object{
				Icon: opt.Something(as.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				}),
				Name:    opt.Something("apple"),
				Summary: opt.Something("banana"),
				URL:     opt.Something("cherry"),
			},
			Expected: []byte(`{`+context+`,"icon":{"url":"http://example.com/img/icon.png"},"name":"apple","summary":"banana","url":"cherry"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual marshaled-jsonld it not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
				t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
