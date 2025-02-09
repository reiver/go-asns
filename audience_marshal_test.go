package asns_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-asns"
)

func TestAudience_marshal(t *testing.T) {

	tests := []struct{
		Value asns.Audience
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Value: asns.Audience{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"name":"apple banana cherry"}`),
		},
		{
			Value: asns.Audience{
				Type: opt.Something("bot"),
			},
			Expected: []byte(`{"type":"bot"}`),
		},
		{
			Value: asns.Audience{
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"url":"http://example.com/basket-bot.html"}`),
		},



		{
			Value: asns.Audience{
				Type: opt.Something("bot"),
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"type":"bot","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Audience{
				Name: opt.Something("apple banana cherry"),
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"name":"apple banana cherry","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Audience{
				Name: opt.Something("apple banana cherry"),
				Type: opt.Something("bot"),
			},
			Expected: []byte(`{"name":"apple banana cherry","type":"bot"}`),
		},



		{
			Value: asns.Audience{
				Name: opt.Something("apple banana cherry"),
				Type: opt.Something("bot"),
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"name":"apple banana cherry","type":"bot","url":"http://example.com/basket-bot.html"}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := asns.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled %T is not what was expected", testNumber, test.Value)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
