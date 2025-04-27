package set

import (
	"slices"
	"strings"
	"testing"
)

func TestValues(t *testing.T) {
	s := Set[int]{1: {}, 2: {}, 3: {}}

	collected := make(map[int]bool)

	iter := Values(s)
	iter(func(val int) bool {
		collected[val] = true
		return true // continue iteration
	})

	for val := range s {
		if !collected[val] {
			t.Errorf("Values() missing value: %d", val)
		}
	}

	if len(collected) != len(s) {
		t.Errorf("Values() collected extra elements: got %d, want %d", len(collected), len(s))
	}
}

func TestValues_EarlyExit(t *testing.T) {
	s := Set[int]{10: {}, 20: {}, 30: {}}

	iter := Values(s)

	count := 0
	iter(func(val int) bool {
		count++
		return false // break after first element
	})

	if count != 1 {
		t.Errorf("Values() did not exit early, yielded %d elements", count)
	}
}

func TestAddAndContains(t *testing.T) {
	s := Set[int]{}
	Add(s, 1, 2, 3)

	if !Contains(s, 1) || !Contains(s, 2) || !Contains(s, 3) {
		t.Errorf("Set should contain added elements")
	}
	if Contains(s, 4) {
		t.Errorf("Set should not contain element not added")
	}
}

func TestDelete(t *testing.T) {
	s := Set[int]{}
	Add(s, 1, 2, 3, 4)
	Delete(s, 2, 4)

	if Contains(s, 2) {
		t.Errorf("Set should not contain deleted element")
	}
}

func TestDeleteSeq(t *testing.T) {
	s := Set[int]{}
	Add(s, 1, 2, 3, 4)

	//>	toDelete := Set[int]{2: {}, 4: {}}
	toDelete := []int{2, 4}
	DeleteSeq(s, slices.Values(toDelete))

	if Contains(s, 2) || Contains(s, 4) {
		t.Errorf("Set should not contain deleted elements")
	}
	if !Contains(s, 1) || !Contains(s, 3) {
		t.Errorf("Set should still contain non-deleted elements")
	}
}

func TestCollect(t *testing.T) {
	input := []int{1, 2, 3, 2, 1}
	seq := func(yield func(int) bool) {
		for _, v := range input {
			yield(v)
		}
	}
	s := Collect(seq)

	if !Contains(s, 1) || !Contains(s, 2) || !Contains(s, 3) {
		t.Errorf("Set should collect all unique elements")
	}
}

func TestMap(t *testing.T) {
	s := Set[int]{1: {}, 2: {}, 3: {}}
	mapped := Map(s, func(x int) int {
		return x * 2
	})

	if !Contains(mapped, 2) || !Contains(mapped, 4) || !Contains(mapped, 6) {
		t.Errorf("Map should apply function to all elements")
	}
}

func TestFilter(t *testing.T) {
	s := Set[int]{1: {}, 2: {}, 3: {}, 4: {}}
	filtered := Filter(s, func(x int) bool {
		return x%2 == 0
	})

	if !Contains(filtered, 2) || !Contains(filtered, 4) {
		t.Errorf("Filter should keep elements matching predicate")
	}
	if Contains(filtered, 1) || Contains(filtered, 3) {
		t.Errorf("Filter should remove elements not matching predicate")
	}
}

func TestUnion(t *testing.T) {
	s1 := Set[int]{1: {}, 2: {}}
	s2 := Set[int]{2: {}, 3: {}}

	union := Union(s1, s2)

	if !Contains(union, 1) || !Contains(union, 2) || !Contains(union, 3) {
		t.Errorf("Union should contain all unique elements from both sets")
	}
}

func TestIntersection(t *testing.T) {
	s1 := Set[int]{1: {}, 2: {}, 3: {}}
	s2 := Set[int]{2: {}, 3: {}, 4: {}}

	intersection := Intersection(s1, s2)

	if Contains(intersection, 1) || !Contains(intersection, 2) || !Contains(intersection, 3) || Contains(intersection, 4) {
		t.Errorf("Union should contain only elements which both sets contaim")
	}
}

func TestSetString(t *testing.T) {
	s := Set[int]{1: {}, 2: {}, 3: {}}

	output := s.String()

	if !strings.HasPrefix(output, "{") || !strings.HasSuffix(output, "}") {
		t.Errorf("Set.String() should start with '{' and end with '}', got: %q", output)
	}

	for _, val := range []string{"1", "2", "3"} {
		if !strings.Contains(output, val) {
			t.Errorf("Set.String() missing value: %s, got: %q", val, output)
		}
	}
}

func TestSetString_EmptySet(t *testing.T) {
	s := Set[int]{}

	output := s.String()

	if output != "{}" {
		t.Errorf("Empty Set.String() should return value: '{}', got '%s'", output)
	}
}
