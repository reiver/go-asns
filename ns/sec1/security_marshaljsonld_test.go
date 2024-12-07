package sec1_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-act"
	"github.com/reiver/go-act/ns/sec1"
)

func TestToot_MarshalJSONLD(t *testing.T) {

	const context string =
	`"@context":[`+
		`"https://w3id.org/security/v1"`+
	`]`

	tests := []struct{
		Value sec1.Security
		Expected []byte
	}{
		{
			Value: sec1.Security{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: sec1.Security{
				PublicKey: opt.Something(
					sec1.PublicKey{
						ID:    opt.Something("http://here.example/~joeblow#main-key"),
					},
				),
			},
			Expected: []byte(`{`+context+`,"publicKey":{"id":"http://here.example/~joeblow#main-key"}}`),
		},
		{
			Value: sec1.Security{
				PublicKey: opt.Something(
					sec1.PublicKey{
						Owner: opt.Something("http://here.example/~joeblow"),
					},
				),
			},
			Expected: []byte(`{`+context+`,"publicKey":{"owner":"http://here.example/~joeblow"}}`),
		},
		{
			Value: sec1.Security{
				PublicKey: opt.Something(
					sec1.PublicKey{
						PublicKeyPem: opt.Something("-----BEGIN PUBLIC KEY-----\nblah blah blah\n-----END PUBLIC KEY-----\n"),
					},
				),
			},
			Expected: []byte(`{`+context+`,"publicKey":{"publicKeyPem":"-----BEGIN PUBLIC KEY-----\nblah blah blah\n-----END PUBLIC KEY-----\n"}}`),
		},



		{
			Value: sec1.Security{
				PublicKey: opt.Something(
					sec1.PublicKey{
						ID:    opt.Something("http://here.example/~joeblow#main-key"),
						Owner: opt.Something("http://here.example/~joeblow"),
						PublicKeyPem: opt.Something("-----BEGIN PUBLIC KEY-----\nblah blah blah\n-----END PUBLIC KEY-----\n"),
					},
				),
			},
			Expected: []byte(`{`+context+`,"publicKey":{"id":"http://here.example/~joeblow#main-key","owner":"http://here.example/~joeblow","publicKeyPem":"-----BEGIN PUBLIC KEY-----\nblah blah blah\n-----END PUBLIC KEY-----\n"}}`),
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
