package asns_test

import (
	"testing"

	"github.com/reiver/go-asns"
)

func TestTypes_MarshalJSON(t *testing.T) {

	tests := []struct{
		Types asns.Types
		Expected string
	}{
		{
			Expected: "null",
		},


		{
			Types: asns.SomeType("apple"),
			Expected:           `"apple"`,
		},
		{
			Types: asns.SomeType("BANANA"),
			Expected:           `"BANANA"`,
		},
		{
			Types: asns.SomeType("Cherry"),
			Expected:           `"Cherry"`,
		},
		{
			Types: asns.SomeType("dAtE"),
			Expected:           `"dAtE"`,
		},



		{
			Types: asns.SomeTypes("apple","BANANA"),
			Expected:           `["apple","BANANA"]`,
		},
		{
			Types: asns.SomeTypes("apple","BANANA","Cherry"),
			Expected:           `["apple","BANANA","Cherry"]`,
		},
		{
			Types: asns.SomeTypes("apple","BANANA","Cherry","dAtE"),
			Expected:           `["apple","BANANA","Cherry","dAtE"]`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Types.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("TYPES: %#v", test.Types)
			continue
		}

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%s", expected)
			t.Logf("ACTUAL:\n%s", actual)
			t.Logf("TYPES: %#v", test.Types)
			continue
		}
	}
}
