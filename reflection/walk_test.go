package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Friend Friend
}

type Friend struct {
	Name string
	Age int
}

func TestWalk(t *testing.T) {
	
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"bob"},
			[]string{"bob"},
		},
		{
			"struct with 2 fields",
			struct {
				Name string
				Friend string
			}{"bob","billy"},
			[]string{"bob","billy"},
		},
		{
			"struct with a nonstring field",
			struct {
				Name string
				Id int
			}{"jo", 1},
			[]string{"jo"},
		},
		{
			"nested structs",
			Person{
				"brian",
				Friend{"lee", 100},
			},
			[]string{"brian", "lee"},
		},
		{
			"pointer to struct",
			&Person{
				"boily",
				Friend{"billiam", 5},
			},
			[]string{"boily", "billiam"},
		},
		{
			"slices",
			[]Friend{
				{"me", 12},
				{"you", 9},
			},
			[]string{"me", "you"},
		},
		{
			"arrays",
			[2]Friend{
				{"beavis", 69},
				{"butthead", 69},
			},
			[]string{"beavis", "butthead"},
		},
	}
	
	for _,test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			want := test.ExpectedCalls
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}

	t.Run("testing maps", func(t *testing.T) {
		testCase := map[string]string {
					"holy": "calamity",
					"scream": "insanity",
		}
		
		var got []string
		walk(testCase, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "calamity")
		assertContains(t, got, "insanity")
	})

	t.Run("testing channels", func(t *testing.T) {
		aChannel := make(chan Friend)

		go func() {
			aChannel <- Friend{"bob", 6}
			aChannel <- Friend{"rob", 7}
			close(aChannel)
		}()

		var got []string
		want := []string{"bob", "rob"}
		
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("testing functions", func(t *testing.T) {
		aFunction := func() (f1, f2 Friend) {
			return Friend{"beep", 3}, Friend{"boop", 1}
		}

		var got []string
		want := []string{"beep", "boop"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _,x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("wanted %+v to contain %q but it did not", haystack, needle)
	}
}