package act_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-act"
	"github.com/reiver/go-act/ns/ap"
	"github.com/reiver/go-act/ns/toot"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Values []any
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Values: []any{
				ap.Actor{},
			},
			Expected: []byte(`{}`),
		},
		{
			Values: []any{
				ap.Actor{
					PreferredUserName: opt.Something("joeblow"),
				},
			},
			Expected: []byte(`{"preferredUsername":"joeblow"}`),
		},
		{
			Values: []any{
				ap.Actor{
					Inbox:  opt.Something("http://social.example/~me/inbox"),
					Outbox: opt.Something("http://social.example/~me/outbox"),
					PreferredUserName: opt.Something("joeblow"),
				},
			},
			Expected: []byte(`{"inbox":"http://social.example/~me/inbox","outbox":"http://social.example/~me/outbox","preferredUsername":"joeblow"}`),
		},



		{
			Values: []any{
				ap.Actor{
					Inbox:  opt.Something("http://social.example/~me/inbox"),
					Outbox: opt.Something("http://social.example/~me/outbox"),
					PreferredUserName: opt.Something("joeblow"),
				},
				toot.Toot{},
			},
			Expected: []byte(
				`{`+
					`"@context":{`+
						`"toot":"http://joinmastodon.org/ns#","attributionDomains":"toot:attributionDomains","blurhash":"toot:blurhash","discoverable":"toot:discoverable","Emoji":"toot:Emoji","featured":"toot:featured","featuredTags":"toot:featuredTags","focalPoint":"toot:focalPoint","IdentityProof":"toot:IdentityProof","indexable":"toot:indexable","memorial":"toot:memorial","votersCount":"toot:votersCount","suspended":"toot:suspended"`+
					`}`+
					`,`+
					`"inbox":"http://social.example/~me/inbox"`+
					`,`+
					`"outbox":"http://social.example/~me/outbox"`+
					`,`+
					`"preferredUsername":"joeblow"`+
				`}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := act.Marshal(test.Values...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an erorr but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled value is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			continue
		}
	}
}
