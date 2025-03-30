package asns_test

import (
	"testing"

	"github.com/reiver/go-asns"
)

func TestTypes_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON []byte
		Expected asns.Types
	}{
		{
			JSON: []byte("null"),
		},


		{
			JSON:           []byte(`"apple"`),
			Expected: asns.SomeType("apple"),
		},
		{
			JSON:           []byte(`"BANANA"`),
			Expected: asns.SomeType("BANANA"),
		},
		{
			JSON:           []byte(`"Cherry"`),
			Expected: asns.SomeType("Cherry"),
		},
		{
			JSON:           []byte(`"dAtE"`),
			Expected: asns.SomeType("dAtE"),
		},



		{
			JSON:           []byte(`["apple","BANANA"]`),
			Expected: asns.SomeTypes("apple","BANANA"),
		},
		{
			JSON:           []byte(`["apple","BANANA","Cherry"]`),
			Expected: asns.SomeTypes("apple","BANANA","Cherry"),
		},
		{
			JSON:           []byte(`["apple","BANANA","Cherry","dAtE"]`),
			Expected: asns.SomeTypes("apple","BANANA","Cherry","dAtE"),
		},
	}

	for testNumber, test := range tests {

		var actual asns.Types

		err := actual.UnmarshalJSON(test.JSON)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}
	}
}
