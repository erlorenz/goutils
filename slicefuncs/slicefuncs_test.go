package slicefuncs_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/erlorenz/goutils/slicefuncs"
)

type person struct {
	name string
	age  age
}

type age int

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	doubled := slicefuncs.Map(nums, func(num int) int {
		return num * 2
	})

	want := []int{2, 4, 6, 8, 10}
	got := doubled
	if !slices.Equal(want, got) {
		t.Errorf("failed equality check: want %v, got %v", want, got)
	}

	people := []person{
		{"Fred", 52},
		{"James", 30},
		{"Jessica", 25},
	}

	namesOnly := slicefuncs.Map(people, func(person person) string {
		return person.name
	})

	want2 := []string{"Fred", "James", "Jessica"}
	got2 := namesOnly

	if !slices.Equal(want2, got2) {
		t.Errorf("failed equality check: want %v, got %v", want, got)
	}
}

func TestFilter(t *testing.T) {
	people := []person{
		{"Fred", 52},
		{"James", 30},
		{"Jessica", 25},
		{"Molly", 60},
	}
	filtered := slicefuncs.Filter(people, func(person person) bool {
		return strings.HasPrefix(person.name, "J")
	})

	want := []person{{"James", 30}, {"Jessica", 25}}
	got := filtered

	if !slices.Equal(want, got) {
		t.Errorf("failed equality check: want %v, got %v", want, got)
	}

	filtered2 := slicefuncs.Filter(people, func(person person) bool {
		return person.age > 55
	})

	want2 := []person{{"Molly", 60}}
	got2 := filtered2

	if !slices.Equal(want2, got2) {
		t.Errorf("failed equality check: want %v, got %v", want2, got2)
	}
}
