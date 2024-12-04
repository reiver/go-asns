package toot_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-act/ns/toot"
)

func TestToot_MarshalJSONLD(t *testing.T) {

	var context string =
	`"@context":{`+
		`"toot":"http://joinmastodon.org/ns#"`+
		`,`+
		`"Curve25519Key":"toot:Curve25519Key"`+
		`,`+
		`"cipherText":"toot:cipherText"`+
		`,`+
		`"claim":"toot:claim"`+
		`,`+
		`"Device":"toot:Device"`+
		`,`+
		`"deviceId":"toot:deviceId"`+
		`,`+
		`"discoverable":"toot:discoverable"`+
		`,`+
		`"Ed25519Key":"toot:Ed25519Key"`+
		`,`+
		`"Ed25519Signature":"toot:Ed25519Signature"`+
		`,`+
		`"Emoji":"toot:Emoji"`+
		`,`+
		`"EncryptedMessage":"toot:EncryptedMessage"`+
		`,`+
		`"featured":"toot:featured"`+
		`,`+
		`"featuredTags":"toot:featuredTags"`+
		`,`+
		`"fingerprintKey":"toot:fingerprintKey"`+
		`,`+
		`"focalPoint":"toot:focalPoint"`+
		`,`+
		`"identityKey":"toot:identityKey"`+
		`,`+
		`"indexable":"toot:indexable"`+
		`,`+
		`"memorial":"toot:memorial"`+
		`,`+
		`"messageFranking":"toot:messageFranking"`+
		`,`+
		`"messageType":"toot:messageType"`+
		`,`+
		`"publicKeyBase64":"toot:publicKeyBase64"`+
		`,`+
		`"suspended":"toot:suspended"`+
	`}`

	tests := []struct{
		Value toot.Toot
		Expected []byte
	}{
		{
			Value: toot.Toot{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: toot.Toot{
				Curve25519Key: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"Curve25519Key":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				CipherText: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"cipherText":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Claim: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"claim":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Device: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"Device":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				DeviceID: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"deviceId":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Discoverable: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"discoverable":true}`),
		},
		{
			Value: toot.Toot{
				Discoverable: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"discoverable":false}`),
		},
		{
			Value: toot.Toot{
				Ed25519Key: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"Ed25519Key":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Ed25519Signature: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"Ed25519Signature":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Emoji: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"Emoji":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				EncryptedMessage: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"EncryptedMessage":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Featured: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"featured":true}`),
		},
		{
			Value: toot.Toot{
				Featured: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"featured":false}`),
		},
		{
			Value: toot.Toot{
				FeaturedTags: []string{"apple","banana","cherry"},
			},
			Expected: []byte(`{`+context+`,"featuredTags":["apple","banana","cherry"]}`),
		},
		{
			Value: toot.Toot{
				FingerPrintKey: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"fingerprintKey":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				FingerPrintKey: opt.Something(""),
			},
			Expected: []byte(`{`+context+`,"fingerprintKey":""}`),
		},
		{
			Value: toot.Toot{
				FocalPoint: []any{12,456},
			},
			Expected: []byte(`{`+context+`,"focalPoint":[12,456]}`),
		},
		{
			Value: toot.Toot{
				IdentityKey: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"identityKey":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Indexable: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"indexable":true}`),
		},
		{
			Value: toot.Toot{
				Indexable: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"indexable":false}`),
		},
		{
			Value: toot.Toot{
				Memorial: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"memorial":true}`),
		},
		{
			Value: toot.Toot{
				Memorial: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"memorial":false}`),
		},
		{
			Value: toot.Toot{
				MessageFranking: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"messageFranking":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				MessageType: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"messageType":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				PublicKeyBase64: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{`+context+`,"publicKeyBase64":"apple banana cherry"}`),
		},
		{
			Value: toot.Toot{
				Suspended: opt.Something(true),
			},
			Expected: []byte(`{`+context+`,"suspended":true}`),
		},
		{
			Value: toot.Toot{
				Suspended: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"suspended":false}`),
		},



		{
			Value: toot.Toot{
				Discoverable: opt.Something(false),
				IdentityKey: opt.Something("apple banana cherry"),
				Indexable: opt.Something(true),
				Memorial: opt.Something(false),
				Suspended: opt.Something(false),
			},
			Expected: []byte(`{`+context+`,"discoverable":false,"identityKey":"apple banana cherry","indexable":true,"memorial":false,"suspended":false}`),
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
				t.Logf("EXPECTED: (%d) %q", len(expected), expected)
				t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
