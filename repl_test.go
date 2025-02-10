package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input string
    expected []string
  }{
    {
      input: "  hello world   ",
      expected: []string{"hello", "world"},
    },
    {
      input: "foo    BAR",
      expected: []string{"foo", "bar"},
    },
    {
      input: "a LONGER strING with SOme   varience ",
      expected: []string{"a", "longer", "string", "with", "some", "varience"},
    },
  }

  for _, c := range cases {
    actual := cleanInput(c.input)
    if len(actual) != len(c.expected) {
      t.Errorf("Length of expected and actual do not match!\nExpected (%v):\n%v\nActual (%v):\n%v",
      len(c.expected), c.expected, len(actual), actual)
    }

    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]

      if word != expectedWord {
        t.Errorf("Expected word is not equal to actual word! %v != %v", word, expectedWord)
      }
    }
  }
}
