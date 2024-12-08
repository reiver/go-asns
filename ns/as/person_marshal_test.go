package as_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-act"
	"github.com/reiver/go-act/ns/as"
)

func TestPerson_Marshal(t *testing.T) {

	const context string =
		`"@context":{`+
			`"as":"https://www.w3.org/ns/activitystreams"`+
			`,`+
			`"alsoKnownAs":"as:alsoKnownAs"`+
			`,`+
			`"attachment":"as:attachment"`+
			`,`+
			`"id":"as:id"`+
			`,`+
			`"image":"as:image"`+
			`,`+
			`"icon":"as:icon"`+
			`,`+
			`"mediaType":"as:mediaType"`+
			`,`+
			`"movedTo":"as:movedTo"`+
			`,`+
			`"name":"as:name"`+
			`,`+
			`"summary":"as:summary"`+
			`,`+
			`"tag":"as:tag"`+
			`,`+
			`"type":"as:type"`+
			`,`+
			`"url":"as:url"`+
		`}`

	tests := []struct{
		Value as.Person
		Expected []byte
	}{
		// 0
		{
			Value: as.Person{},
			Expected: []byte(`{`+context+`,"type":"Person"}`),
		},



		// 1
		{
			Value: as.Person{
				AlsoKnownAs: nil,
			},
			Expected: []byte(`{`+context+`,"type":"Person"}`),
		},
		// 2
		{
			Value: as.Person{
				AlsoKnownAs: []string{},
			},
			Expected: []byte(`{`+context+`,"type":"Person"}`),
		},
		// 3
		{
			Value: as.Person{
				AlsoKnownAs: []string{"once"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once"],"type":"Person"}`),
		},
		// 4
		{
			Value: as.Person{
				AlsoKnownAs: []string{"once","twice"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once","twice"],"type":"Person"}`),
		},
		// 5
		{
			Value: as.Person{
				AlsoKnownAs: []string{"once","twice","thrice"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once","twice","thrice"],"type":"Person"}`),
		},
		// 6
		{
			Value: as.Person{
				AlsoKnownAs: []string{"once","twice","thrice","fource"},
			},
			Expected: []byte(`{`+context+`,"alsoKnownAs":["once","twice","thrice","fource"],"type":"Person"}`),
		},



		// 7
		{
			Value: as.Person{
				Icon: as.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				},
			},
			Expected: []byte(`{`+context+`,"icon":{"url":"http://example.com/img/icon.png"},"type":"Person"}`),
		},
		// 8
		{
			Value: as.Person{
				Icon: as.Icon{
					Height: opt.Something(uint64(123)),
					Name: opt.Something("apple banana cherry"),
					Type: opt.Something("image/png"),
					URL: opt.Something("http://example.com/img/icon.png"),
					Width: opt.Something(uint64(456)),
				},
			},
			Expected: []byte(`{`+context+`,"icon":{"height":123,"name":"apple banana cherry","type":"image/png","url":"http://example.com/img/icon.png","width":456},"type":"Person"}`),
		},
		// 9
		{
			Value: as.Person{
				MovedTo: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"movedTo":"apple banana cherry","type":"Person"}`),
		},
		// 10
		{
			Value: as.Person{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"name":"apple banana cherry","type":"Person"}`),
		},
		// 11
		{
			Value: as.Person{
				Summary: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"summary":"apple banana cherry","type":"Person"}`),
		},
		// 12
		{
			Value: as.Person{
				URL: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"type":"Person","url":"apple banana cherry"}`),
		},



		// 13
		{
			Value: as.Person{
				Icon: as.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				},
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
					`,`+
					`"type":"Person"`+
				`}`),
		},



		// 14
		{
			Value: as.Person{
				Name:    opt.Something("apple"),
				Summary: opt.Something("banana"),
				URL:     opt.Something("cherry"),
			},
			Expected: []byte(`{`+context+`,"name":"apple","summary":"banana","type":"Person","url":"cherry"}`),
		},
		// 15
		{
			Value: as.Person{
				Icon: as.Icon{
					URL: opt.Something("http://example.com/img/icon.png"),
				},
				Name:    opt.Something("apple"),
				Summary: opt.Something("banana"),
				URL:     opt.Something("cherry"),
			},
			Expected: []byte(`{`+context+`,"icon":{"url":"http://example.com/img/icon.png"},"name":"apple","summary":"banana","type":"Person","url":"cherry"}`),
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
