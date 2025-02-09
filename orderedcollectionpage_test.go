package asns_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-asns"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

func TestOrderedCollectionPage(t *testing.T) {

	const context string = `"@context":{"as":"https://www.w3.org/ns/activitystreams","audience":"as:audience","generator":"as:generator","height":"as:height","icon":"as:icon","id":"as:id","image":"as:image","location":"as:location","mediaType":"as:mediaType","name":"as:name","next":"as:next","orderedItems":"as:orderedItems","partOf":"as:partOf","prev":"as:prev","published":"as:published","startIndex":"as:startIndex","summary":"as:summary","type":"as:type","updated":"as:updated","url":"as:url","width":"as:width"}`


	tests := []struct{
		Value asns.OrderedCollectionPage
		Expected []byte
	}{
		{
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},



		{
			Value: asns.OrderedCollectionPage{
				ID: opt.Something("acct:joeblow@example.com/outbox/page/?num=4"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"id":"acct:joeblow@example.com/outbox/page/?num=4"`+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				Name: opt.Something("Outbox Page 4"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"name":"Outbox Page 4"`+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				Next: opt.Something("acct:joeblow@example.com/outbox/page/?num=5"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"next":"acct:joeblow@example.com/outbox/page/?num=5"`+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				PartOf: opt.Something("acct:joeblow@example.com/outbox/"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"partOf":"acct:joeblow@example.com/outbox/"`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				Prev: opt.Something("acct:joeblow@example.com/outbox/page/?num=3"),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"prev":"acct:joeblow@example.com/outbox/page/?num=3"`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				StartIndex: opt.Something(uint64(1)),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"startIndex":1`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				Summary: opt.Something(`Page 4 or the Outbox of <a href="https://example.com/~joeblow">@joeblow@example.com</a>`),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"summary":"Page 4 or the Outbox of \u003ca href=\"https://example.com/~joeblow\"\u003e@joeblow@example.com\u003c/a\u003e"`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},



		{
			Value: asns.OrderedCollectionPage{
				OrderedItems: []any{},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				OrderedItems: []any{
					asns.Announce{
						Object: opt.Something("http://host.example/~janedoe/statuses/0000000001"),
					},
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[`+
						`{"object":"http://host.example/~janedoe/statuses/0000000001","type":"Announce"}`+
					`]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				OrderedItems: []any{
					asns.Announce{
						Object: opt.Something("http://host.example/~johndoe/statuses/7050301090"),
					},
					asns.Announce{
						Object: opt.Something("http://host.example/~janedoe/statuses/0000000001"),
					},
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"orderedItems":[`+
						`{"object":"http://host.example/~johndoe/statuses/7050301090","type":"Announce"}`+
						`,`+
						`{"object":"http://host.example/~janedoe/statuses/0000000001","type":"Announce"}`+
					`]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},



		{
			Value: asns.OrderedCollectionPage{
				Icon: asns.Icon{
					URL: opt.Something("https://img.example/icons/something.png"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"icon":{"type":"Image","url":"https://img.example/icons/something.png"}`+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				Image: asns.Image{
					URL: opt.Something("https://img.example/images/something.png"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"image":{"type":"Image","url":"https://img.example/images/something.png"}`+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},









		{
			Value: asns.OrderedCollectionPage{
				ID: opt.Something("acct:joeblow@example.com/outbox/page/?num=4"),
				Icon: asns.Image{
					URL: opt.Something("https://img.example/icons/something.png"),
				},
				Image: asns.Image{
					URL: opt.Something("https://img.example/images/something.png"),
				},
				Name: opt.Something("Outbox Page 4"),
				Next: opt.Something("acct:joeblow@example.com/outbox/page/?num=5"),
				PartOf: opt.Something("acct:joeblow@example.com/outbox/"),
				Prev: opt.Something("acct:joeblow@example.com/outbox/page/?num=3"),
				Summary: opt.Something(`Page 4 or the Outbox of <a href="https://example.com/~joeblow">@joeblow@example.com</a>`),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"icon":{"type":"Image","url":"https://img.example/icons/something.png"}`+
					`,`+
					`"id":"acct:joeblow@example.com/outbox/page/?num=4"`+
					`,`+
					`"image":{"type":"Image","url":"https://img.example/images/something.png"}`+
					`,`+
					`"name":"Outbox Page 4"`+
					`,`+
					`"next":"acct:joeblow@example.com/outbox/page/?num=5"`+
					`,`+
					`"orderedItems":[]`+
					`,`+
					`"partOf":"acct:joeblow@example.com/outbox/"`+
					`,`+
					`"prev":"acct:joeblow@example.com/outbox/page/?num=3"`+
					`,`+
					`"summary":"Page 4 or the Outbox of \u003ca href=\"https://example.com/~joeblow\"\u003e@joeblow@example.com\u003c/a\u003e"`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
		},
		{
			Value: asns.OrderedCollectionPage{
				ID: opt.Something("acct:joeblow@example.com/outbox/page/?num=4"),
				Icon: asns.Image{
					URL: opt.Something("https://img.example/icons/something.png"),
				},
				Image: asns.Image{
					URL: opt.Something("https://img.example/images/something.png"),
				},
				Name: opt.Something("Outbox Page 4"),
				Next: opt.Something("acct:joeblow@example.com/outbox/page/?num=5"),
				OrderedItems: []any{
					asns.Announce{
						Object: opt.Something("http://host.example/~johndoe/statuses/7050301090"),
					},
					asns.Announce{
						Object: opt.Something("http://host.example/~janedoe/statuses/0000000001"),
					},
				},
				PartOf: opt.Something("acct:joeblow@example.com/outbox/"),
				Prev: opt.Something("acct:joeblow@example.com/outbox/page/?num=3"),
				Summary: opt.Something(`Page 4 or the Outbox of <a href="https://example.com/~joeblow">@joeblow@example.com</a>`),
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"icon":{"type":"Image","url":"https://img.example/icons/something.png"}`+
					`,`+
					`"id":"acct:joeblow@example.com/outbox/page/?num=4"`+
					`,`+
					`"image":{"type":"Image","url":"https://img.example/images/something.png"}`+
					`,`+
					`"name":"Outbox Page 4"`+
					`,`+
					`"next":"acct:joeblow@example.com/outbox/page/?num=5"`+
					`,`+
					`"orderedItems":[`+
						`{"object":"http://host.example/~johndoe/statuses/7050301090","type":"Announce"}`+
						`,`+
						`{"object":"http://host.example/~janedoe/statuses/0000000001","type":"Announce"}`+
					`]`+
					`,`+
					`"partOf":"acct:joeblow@example.com/outbox/"`+
					`,`+
					`"prev":"acct:joeblow@example.com/outbox/page/?num=3"`+
					`,`+
					`"summary":"Page 4 or the Outbox of \u003ca href=\"https://example.com/~joeblow\"\u003e@joeblow@example.com\u003c/a\u003e"`+
					`,`+
					`"type":"OrderedCollectionPage"`+
				`}`,
			),
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

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled OrderedCollectionPage json-ld is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
