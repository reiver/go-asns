package asns_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonpp"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-asns"
)

func TestPerson_marshal(t *testing.T) {

	const context string =
		`"@context":{`+
			`"as":"https://www.w3.org/ns/activitystreams"`+
			`,`+
			`"alsoKnownAs":"as:alsoKnownAs"`+
			`,`+
			`"attachment":"as:attachment"`+
			`,`+
			`"audience":"as:audience"`+
			`,`+
			`"generator":"as:generator"`+
			`,`+
			`"height":"as:height"`+
			`,`+
			`"icon":"as:icon"`+
			`,`+
			`"id":"as:id"`+
			`,`+
			`"image":"as:image"`+
			`,`+
			`"location":"as:location"`+
			`,`+
			`"mediaType":"as:mediaType"`+
			`,`+
			`"movedTo":"as:movedTo"`+
			`,`+
			`"name":"as:name"`+
			`,`+
			`"published":"as:published"`+
			`,`+
			`"summary":"as:summary"`+
			`,`+
			`"tag":"as:tag"`+
			`,`+
			`"type":"as:type"`+
			`,`+
			`"updated":"as:updated"`+
			`,`+
			`"url":"as:url"`+
			`,`+
			`"width":"as:width"`+
		`}`

	tests := []struct{
		Value asns.Person
		Expected []byte
	}{
		// 0
		{
			Value: asns.Person{},
			Expected: []byte(`{`+context+`,"type":"Person"}`),
		},



		// 1
		{
			Value: asns.Person{
				AlsoKnownAs: nil,
			},
			Expected: []byte(`{`+context+`,"type":"Person"}`),
		},
		// 2
		{
			Value: asns.Person{
				AlsoKnownAs: []string{},
			},
			Expected: []byte(`{`+context+`,"type":"Person"}`),
		},
		// 3
		{
			Value: asns.Person{
				AlsoKnownAs: []string{"once"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once"],"type":"Person"}`),
		},
		// 4
		{
			Value: asns.Person{
				AlsoKnownAs: []string{"once","twice"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once","twice"],"type":"Person"}`),
		},
		// 5
		{
			Value: asns.Person{
				AlsoKnownAs: []string{"once","twice","thrice"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once","twice","thrice"],"type":"Person"}`),
		},
		// 6
		{
			Value: asns.Person{
				AlsoKnownAs: []string{"once","twice","thrice","fource"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once","twice","thrice","fource"],"type":"Person"}`),
		},



		// 7
		{
			Value: asns.Person{
				Icon: asns.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				},
			},
			Expected: []byte(`{`+context+`,"icon":{"type":"Image","url":"http://example.com/img/icon.png"},"type":"Person"}`),
		},
		// 8
		{
			Value: asns.Person{
				Icon: asns.Icon{
					Height: opt.Something(uint64(123)),
					MediaType: opt.Something("image/png"),
					Name: opt.Something("apple banana cherry"),
					URL: opt.Something("http://example.com/img/icon.png"),
					Width: opt.Something(uint64(456)),
				},
			},
			Expected: []byte(`{`+context+`,"icon":{"height":123,"mediaType":"image/png","name":"apple banana cherry","type":"Image","url":"http://example.com/img/icon.png","width":456},"type":"Person"}`),
		},
		// 9
		{
			Value: asns.Person{
				MovedTo: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"movedTo":"apple banana cherry","type":"Person"}`),
		},
		// 10
		{
			Value: asns.Person{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"name":"apple banana cherry","type":"Person"}`),
		},
		// 11
		{
			Value: asns.Person{
				Summary: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"summary":"apple banana cherry","type":"Person"}`),
		},
		// 12
		{
			Value: asns.Person{
				URL: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"type":"Person","url":"apple banana cherry"}`),
		},



		// 13
		{
			Value: asns.Person{
				Icon: asns.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				},
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"icon":{`+
						`"type":"Image"`+
						`,`+
						`"url":"http://example.com/img/icon.png"`+
					`}`+
					`,`+
					`"name":"apple banana cherry"`+
					`,`+
					`"type":"Person"`+
				`}`),
		},



		// 14
		{
			Value: asns.Person{
				Name:    opt.Something("apple"),
				Summary: opt.Something("banana"),
				URL:     opt.Something("cherry"),
			},
			Expected: []byte(`{`+context+`,"name":"apple","summary":"banana","type":"Person","url":"cherry"}`),
		},
		// 15
		{
			Value: asns.Person{
				Icon: asns.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				},
				Name:    opt.Something("apple"),
				Summary: opt.Something("banana"),
				URL:     opt.Something("cherry"),
			},
			Expected: []byte(`{`+context+`,"icon":{"type":"Image","url":"http://example.com/img/icon.png"},"name":"apple","summary":"banana","type":"Person","url":"cherry"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := asns.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual marshaled it not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d)\n%s", len(expected), jsonpp.SPrettyPrint(expected))
				t.Logf("ACTUAL:   (%d)\n%s", len(actual), jsonpp.SPrettyPrint(actual))
				t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
				t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
