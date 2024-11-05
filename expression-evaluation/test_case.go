package expressionevaluation

import "fmt"

var tests = []struct {
	name       string
	input      Input
	expression string
}{
	{
		name: "Simple", input: Input{ID: 2, Name: "Jim Kelly Field", City: "Ait Ali", Country: "Algeria", Currency: "Dinar"},
		expression: `Currency == "Dinar"`,
	},
	{
		name: "Complex", input: Input{ID: 2, Name: "Jim Kelly Field", City: "Ait Ali", Country: "Algeria", Currency: "Dinar"},
		expression: `(Country == "Algeria" || ID == 2) && (Country == "Jim Kelly Field" || City == "Ait Ali") && Currency == "Dinar"`,
	},
}

func testName(prefix string, function string) string {
	return fmt.Sprintf("%s/%s", prefix, function)
}
