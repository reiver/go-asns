package ap_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-act/ns/ap"
)

func TestActor_MarshalJSONLD(t *testing.T) {

	tests := []struct{
		Value ap.Actor
		Expected []byte
	}{
		{
			Value: ap.Actor{},
			Expected: []byte(`{}`),
		},



		{
			Value: ap.Actor{
				EndPoints: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"endpoints":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Following: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"following":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Followers: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"followers":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Inbox: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"inbox":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Liked: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"liked":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Shares: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"shares":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				OauthAuthorizationEndPoint: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"oauthAuthorizationEndpoint":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				OauthTokenEndPoint: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"oauthTokenEndpoint":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Outbox: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"outbox":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				PreferredUserName: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"preferredUsername":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				ProvideClientKey: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"provideClientKey":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				ProxyURL: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"proxyUrl":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				SharedInbox: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"sharedInbox":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				SignClientKey: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"signClientKey":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Source: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"source":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				Streams: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"streams":"apple banana cherry"}`),
		},
		{
			Value: ap.Actor{
				UploadMedia: opt.Something("apple banana cherry"),
			},
			Expected: []byte(`{"uploadMedia":"apple banana cherry"}`),
		},



		{
			Value: ap.Actor{
				Followers:         opt.Something("https://mastonaut.example/users/joeblow/followers"),
				Following:         opt.Something("https://mastonaut.example/users/joeblow/following"),
				Inbox:             opt.Something("https://mastonaut.example/users/joeblow/inbox"),
				Outbox:            opt.Something("https://mastonaut.example/users/joeblow/outbox"),
				PreferredUserName: opt.Something("joeblow"),
			},
			Expected: []byte(`{"following":"https://mastonaut.example/users/joeblow/following","followers":"https://mastonaut.example/users/joeblow/followers","inbox":"https://mastonaut.example/users/joeblow/inbox","outbox":"https://mastonaut.example/users/joeblow/outbox","preferredUsername":"joeblow"}`),
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
