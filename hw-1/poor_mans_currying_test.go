package main

import (
	"github.com/gsamokovarov/assert"
	"testing"
)

func TestRepeaterSample(t *testing.T) {
	actual := Repeater("foo", ":")(3)
	expected := "foo:foo:foo"

	if expected != actual {
		t.Errorf("Expected `%s` but got `%s`", expected, actual)
	}
}

func TestGeneratorSample(t *testing.T) {
	counter := Generator(
		func(v int) int { return v + 1 },
		0,
	)

	var counterResults [4]int
	for i := 0; i < 4; i++ {
		counterResults[i] = counter()
	}

	actual := counterResults
	expected := [4]int{0, 1, 2, 3}

	if expected != actual {
		t.Errorf("Expected `%d` but got `%d`", expected, actual)
	}

	power := Generator(
		func(v int) int { return v * v },
		2,
	)

	var powerResults [4]int
	for i := 0; i < 4; i++ {
		powerResults[i] = power()
	}

	actual = powerResults
	expected = [4]int{2, 4, 16, 256}

	if expected != actual {
		t.Errorf("Expected `%d` but got `%d`", expected, actual)
	}
}

func TestMapReducerSample(t *testing.T) {
	powerSum := MapReducer(
		func(v int) int { return v * v },
		func(a, v int) int { return a + v },
		0,
	)

	actual := powerSum(1, 2, 3, 4)
	expected := 30

	if expected != actual {
		t.Errorf("Expected `%d` but got `%d`", expected, actual)
	}
}

func TestRepeater(t *testing.T) {
	tests := []struct {
		desc     string
		s        string
		sep      string
		times    int
		expected string
	}{
		{
			desc:     "Normal word with normal separator, but a negative number of times",
			s:        "foo",
			sep:      ":",
			times:    -5,
			expected: "",
		},
		{
			desc:     "Normal word with normal separator a normal number of times",
			s:        "foo",
			sep:      ":",
			times:    3,
			expected: "foo:foo:foo",
		},
		{
			desc:     "Normal word with normal separator, but just one time",
			s:        "foo",
			sep:      ":",
			times:    1,
			expected: "foo",
		},
		{
			desc:     "Normal word that normal separator, but zero times",
			s:        "foo",
			sep:      ":",
			times:    0,
			expected: "",
		},
		{
			desc:     "Normal word with empty separator, several times",
			s:        "foo",
			sep:      "",
			times:    3,
			expected: "foofoofoo",
		},
		{
			desc:     "Normal word with newline separator, several times",
			s:        "foo",
			sep:      "\n",
			times:    3,
			expected: "foo\nfoo\nfoo",
		},
		{
			desc:     "Empty word with newline separator, several times",
			s:        "",
			sep:      "\n",
			times:    3,
			expected: "\n\n",
		},
		{
			desc:     "Newline word with newline separator, several times",
			s:        "\n",
			sep:      "\n",
			times:    3,
			expected: "\n\n\n\n\n",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := Repeater(test.s, test.sep)(test.times)
			expected := test.expected
			assert.Equal(t, expected, actual)
		})
	}
}

func TestGenerator(t *testing.T) {
	t.Run("Example from course works", func(t *testing.T) {
		counter := Generator(
			func(v int) int { return v + 1 },
			0,
		)
		sqrt := Generator(
			func(v int) int { return v * v },
			2,
		)

		assert.Equal(t, 0, counter())
		assert.Equal(t, 1, counter())
		assert.Equal(t, 2, sqrt())
		assert.Equal(t, 4, sqrt())
		assert.Equal(t, 2, counter())
		assert.Equal(t, 16, sqrt())
		assert.Equal(t, 3, counter())
		assert.Equal(t, 256, sqrt())
	})
}

func TestMapReducer(t *testing.T) {
	t.Run("Example from course works", func(t *testing.T) {
		sqrtSum := MapReducer(
			func(v int) int { return v * v },
			func(a, v int) int { return a + v },
			0,
		)

		assert.Equal(t, 30, sqrtSum(1, 2, 3, 4))
	})

}
