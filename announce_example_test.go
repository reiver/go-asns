package act_test

import (
	"fmt"

	"github.com/reiver/go-opt"
	"github.com/reiver/go-jsonld"

	"github.com/reiver/go-act"
)

func ExampleAnnounce() {
	var announcement = act.Announce {
		ID: opt.Something("https://mastodon.social/users/reiver/statuses/113962917187191012/activity"),
		Actor:     opt.Something("acct:reiver@mastodon.social"),
		Published: opt.Something("2025-02-07T14:56:43Z"),
		To: []string{"https://www.w3.org/ns/activitystreams#Public"},
		Object: opt.Something("https://example.com/note/0xf1938e3a1a6f4ac790fdf66d1b167858"),
	}

	bytes, err := jsonld.Marshal(announcement)
	if nil != err {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("%s", bytes)
}
