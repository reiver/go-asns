package asns_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-asns"
)

func TestLocation_marshal(t *testing.T) {

	tests := []struct{
		Value asns.Location
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Value: asns.Location{
				Altitude: opt.Something("10.02"),
			},
			Expected: []byte(`{"altitude":10.02}`),
		},
		{
			Value: asns.Location{
				Latitude: opt.Something("9.35"),
			},
			Expected: []byte(`{"latitude":9.35}`),
		},
		{
			Value: asns.Location{
				Longitude: opt.Something("20.14"),
			},
			Expected: []byte(`{"longitude":20.14}`),
		},
		{
			Value: asns.Location{
				Name: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"name":"apple banana cherry"}`),
		},
		{
			Value: asns.Location{
				Type: opt.Something("bot"),
			},
			Expected: []byte(`{"type":"bot"}`),
		},
		{
			Value: asns.Location{
				Units: opt.Something("mm"),
			},
			Expected: []byte(`{"units":"mm"}`),
		},
		{
			Value: asns.Location{
				URL:  opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"url":"http://example.com/basket-bot.html"}`),
		},









		{
			Value: asns.Location{
				Latitude:  opt.Something("9.35"),
				Longitude: opt.Something("20.14"),
				Name:      opt.Something("apple banana cherry"),
				Type:      opt.Something("bot"),
				Units:     opt.Something("mm"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"latitude":9.35,"longitude":20.14,"name":"apple banana cherry","type":"bot","units":"mm","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Longitude: opt.Something("20.14"),
				Name:      opt.Something("apple banana cherry"),
				Type:      opt.Something("bot"),
				Units:     opt.Something("mm"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"altitude":10.02,"longitude":20.14,"name":"apple banana cherry","type":"bot","units":"mm","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Latitude:  opt.Something("9.35"),
				Name:      opt.Something("apple banana cherry"),
				Type:      opt.Something("bot"),
				Units:     opt.Something("mm"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"altitude":10.02,"latitude":9.35,"name":"apple banana cherry","type":"bot","units":"mm","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Latitude:  opt.Something("9.35"),
				Longitude: opt.Something("20.14"),
				Type:      opt.Something("bot"),
				Units:     opt.Something("mm"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"altitude":10.02,"latitude":9.35,"longitude":20.14,"type":"bot","units":"mm","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Latitude:  opt.Something("9.35"),
				Longitude: opt.Something("20.14"),
				Name:      opt.Something("apple banana cherry"),
				Units:     opt.Something("mm"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"altitude":10.02,"latitude":9.35,"longitude":20.14,"name":"apple banana cherry","units":"mm","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Latitude:  opt.Something("9.35"),
				Longitude: opt.Something("20.14"),
				Name:      opt.Something("apple banana cherry"),
				Type:      opt.Something("bot"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"altitude":10.02,"latitude":9.35,"longitude":20.14,"name":"apple banana cherry","type":"bot","url":"http://example.com/basket-bot.html"}`),
		},
		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Latitude:  opt.Something("9.35"),
				Longitude: opt.Something("20.14"),
				Name:      opt.Something("apple banana cherry"),
				Type:      opt.Something("bot"),
				Units:     opt.Something("mm"),
			},
			Expected: []byte(`{"altitude":10.02,"latitude":9.35,"longitude":20.14,"name":"apple banana cherry","type":"bot","units":"mm"}`),
		},









		{
			Value: asns.Location{
				Altitude:  opt.Something("10.02"),
				Latitude:  opt.Something("9.35"),
				Longitude: opt.Something("20.14"),
				Name:      opt.Something("apple banana cherry"),
				Type:      opt.Something("bot"),
				Units:     opt.Something("mm"),
				URL:       opt.Something("http://example.com/basket-bot.html"),
			},
			Expected: []byte(`{"altitude":10.02,"latitude":9.35,"longitude":20.14,"name":"apple banana cherry","type":"bot","units":"mm","url":"http://example.com/basket-bot.html"}`),
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
