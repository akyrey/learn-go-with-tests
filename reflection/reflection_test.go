package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Dario"},
			[]string{"Dario"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Dario", "London"},
			[]string{"Dario", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Dario", 33},
			[]string{"Dario"},
		},
		{
			"nested fields",
			Person{"Dario", Profile{"London", 33}},
			[]string{"Dario", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Dario",
				Profile{"London", 33},
			},
			[]string{"Dario", "London"},
		},
		{
			"slices",
			[]Profile{
				{"London", 33},
				{"Reykjavík", 34},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
				{"London", 33},
				{"Reykjavík", 34},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"Berlin", 33}
			aChannel <- Profile{"Katowice", 34}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"Berlin", 33}, Profile{"Katowice", 34}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}
