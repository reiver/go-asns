package asns_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-asns"
	"github.com/reiver/go-asns/ns/sec1"
	toot "github.com/reiver/go-tootns"
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
				asns.Actor{},
			},
			Expected: []byte(`{}`),
		},
		{
			Values: []any{
				asns.Actor{
					PreferredUserName: opt.Something("joeblow"),
				},
			},
			Expected: []byte(`{"preferredUsername":"joeblow"}`),
		},
		{
			Values: []any{
				asns.Actor{
					Inbox:  opt.Something("http://social.example/~me/inbox"),
					Outbox: opt.Something("http://social.example/~me/outbox"),
					PreferredUserName: opt.Something("joeblow"),
				},
			},
			Expected: []byte(`{"inbox":"http://social.example/~me/inbox","outbox":"http://social.example/~me/outbox","preferredUsername":"joeblow"}`),
		},



		{
			Values: []any{
				asns.Actor{
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
		{
			Values: []any{
				asns.Actor{
					Inbox:  opt.Something("http://social.example/~me/inbox"),
					Outbox: opt.Something("http://social.example/~me/outbox"),
					PreferredUserName: opt.Something("joeblow"),
				},
				toot.Toot{
					Discoverable: opt.Something(true),
					Indexable:    opt.Something(true),
				},
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
					`,`+
					`"discoverable":true`+
					`,`+
					`"indexable":true`+
				`}`),
		},



		{
			Values: []any{
				asns.Actor{
					Inbox:  opt.Something("http://social.example/~me/inbox"),
					Outbox: opt.Something("http://social.example/~me/outbox"),
					PreferredUserName: opt.Something("joeblow"),
				},
				toot.Toot{
					Discoverable: opt.Something(true),
					Indexable:    opt.Something(true),
				},
				sec1.Security{},
			},
			Expected: []byte(
				`{`+
					`"@context":[`+
						`"https://w3id.org/security/v1"`+
						`,`+
						`{`+
							`"toot":"http://joinmastodon.org/ns#","attributionDomains":"toot:attributionDomains","blurhash":"toot:blurhash","discoverable":"toot:discoverable","Emoji":"toot:Emoji","featured":"toot:featured","featuredTags":"toot:featuredTags","focalPoint":"toot:focalPoint","IdentityProof":"toot:IdentityProof","indexable":"toot:indexable","memorial":"toot:memorial","votersCount":"toot:votersCount","suspended":"toot:suspended"`+
						`}`+
					`]`+
					`,`+
					`"inbox":"http://social.example/~me/inbox"`+
					`,`+
					`"outbox":"http://social.example/~me/outbox"`+
					`,`+
					`"preferredUsername":"joeblow"`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"indexable":true`+
				`}`),
		},
		{
			Values: []any{
				asns.Actor{
					Inbox:  opt.Something("http://social.example/~me/inbox"),
					Outbox: opt.Something("http://social.example/~me/outbox"),
					PreferredUserName: opt.Something("joeblow"),
				},
				toot.Toot{
					Discoverable: opt.Something(true),
					Indexable:    opt.Something(true),
				},
				sec1.Security{
					PublicKey: opt.Something(sec1.PublicKey{
						ID: opt.Something("http://social.example/~me#main-key"),
						Owner: opt.Something("http://social.example/~me"),
						PublicKeyPem: opt.Something("-----BEGIN PUBLIC KEY-----\nblah blah blah\n-----END PUBLIC KEY-----\n"),
					}),
				},
			},
			Expected: []byte(
				`{`+
					`"@context":[`+
						`"https://w3id.org/security/v1"`+
						`,`+
						`{`+
							`"toot":"http://joinmastodon.org/ns#","attributionDomains":"toot:attributionDomains","blurhash":"toot:blurhash","discoverable":"toot:discoverable","Emoji":"toot:Emoji","featured":"toot:featured","featuredTags":"toot:featuredTags","focalPoint":"toot:focalPoint","IdentityProof":"toot:IdentityProof","indexable":"toot:indexable","memorial":"toot:memorial","votersCount":"toot:votersCount","suspended":"toot:suspended"`+
						`}`+
					`]`+
					`,`+
					`"inbox":"http://social.example/~me/inbox"`+
					`,`+
					`"outbox":"http://social.example/~me/outbox"`+
					`,`+
					`"preferredUsername":"joeblow"`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"indexable":true`+
					`,`+
					`"publicKey":{`+
						`"id":"http://social.example/~me#main-key"`+
						`,`+
						`"owner":"http://social.example/~me"`+
						`,`+
						`"publicKeyPem":"-----BEGIN PUBLIC KEY-----\nblah blah blah\n-----END PUBLIC KEY-----\n"`+
					`}`+
				`}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := asns.Marshal(test.Values...)
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
