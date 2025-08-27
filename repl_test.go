package main
import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "Hello         Pikachu",
			expected: []string{"hello", "pikachu"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("len(actual) != len(expected)")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words are different in index %d. Actual: %v, expected: %v", i, actual, c.expected)
			}
		}
	}
}
