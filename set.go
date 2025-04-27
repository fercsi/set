package set

import (
	"fmt"
	"iter"
	"strings"
)

// Set is a generic set of comparable elements.
type Set[E comparable] map[E]struct{}

// Values returns a sequence iterator over the elements of the set.
func Values[S Set[E], E comparable](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for val := range s {
			if !yield(val) {
				return
			}
		}
	}
}

// Add inserts one or more elements into the set.
func Add[S Set[E], E comparable](s S, vals ...E) {
	for _, val := range vals {
		s[val] = struct{}{}
	}
}

// Contains reports whether the set contains the specified element.
func Contains[S Set[E], E comparable](s S, val E) bool {
	_, ok := s[val]
	return ok
}

// Delete removes one or more elements from the set.
func Delete[S Set[E], E comparable](s S, vals ...E) {
	for _, val := range vals {
		delete(s, val)
	}
}

// DeleteSeq removes all elements in the sequence from the set.
func DeleteSeq[S Set[E], E comparable](s S, seq iter.Seq[E]) {
	for val := range seq {
		delete(s, val)
	}
}

// Collect creates a set from a sequence of elements.
func Collect[E comparable](seq iter.Seq[E]) Set[E] {
	s := Set[E]{}
	for val := range seq {
		Add(s, val)
	}
	return s
}

// Map applies a function to each element of the set and returns a new set with the results.
func Map[S Set[E], E comparable, R comparable](s S, f func(E) R) Set[R] {
	t := Set[R]{}
	for val := range Values(s) {
		Add(t, f(val))
	}
	return t
}

// Filter returns a set containing elements for which the provided function returns true.
func Filter[S Set[E], E comparable](s S, f func(E) bool) Set[E] {
	t := Set[E]{}
	for val := range Values(s) {
		if f(val) {
			Add(t, val)
		}
	}
	return t
}

// Union returns a set containing all elements from both sets.
func Union[S Set[E], E comparable](s1, s2 S) Set[E] {
	s := Set[E]{}
	for val := range s1 {
		Add(s, val)
	}
	for val := range s2 {
		Add(s, val)
	}
	return s
}

// Intersection returns a set containing only the elements present in both sets.
func Intersection[S Set[E], E comparable](s1, s2 S) Set[E] {
	s := Set[E]{}
	for val := range s1 {
		if Contains(s2, val) {
			Add(s, val)
		}
	}
	return s
}

// String returns a string representation of the set.
//
// This method allows the set to be used with fmt.Print, fmt.Println, and other fmt functions that
// expect a string or any value that implements the Stringer interface. 
//
// Example usage:
//     s := Set[int]{1, 2, 3}
//     fmt.Println(s) // Output: {1 2 3}
func (s Set[E]) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	first := true
	for val := range s {
		if !first {
			sb.WriteString(" ")
		}
		first = false
		sb.WriteString(fmt.Sprint(val))
	}
	sb.WriteString("}")
	return sb.String()
}
