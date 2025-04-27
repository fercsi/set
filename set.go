package set

import (
	"fmt"
	"iter"
	"strings"
)

type Set[E comparable] map[E]struct{}

func Values[S Set[E], E comparable](s S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for val := range s {
			if !yield(val) {
				return
			}
		}
	}
}

// Add inserts the values into the given set `s`.
func Add[S Set[E], E comparable](s S, vals ...E) {
	for _, val := range vals {
		s[val] = struct{}{}
	}
}

func Contains[S Set[E], E comparable](s S, val E) bool {
	_, ok := s[val]
	return ok
}

func Delete[S Set[E], E comparable](s S, vals ...E) {
	for _, val := range vals {
		delete(s, val)
	}
}

func DeleteSeq[S Set[E], E comparable](s S, seq iter.Seq[E]) {
	for val := range seq {
		delete(s, val)
	}
}

func Collect[E comparable](seq iter.Seq[E]) Set[E] {
	s := Set[E]{}
	for val := range seq {
		Add(s, val)
	}
	return s
}

func Map[S Set[E], E comparable, R comparable](s S, f func(E) R) Set[R] {
	t := Set[R]{}
	for val := range Values(s) {
		Add(t, f(val))
	}
	return t
}

func Filter[S Set[E], E comparable](s S, f func(E) bool) Set[E] {
	t := Set[E]{}
	for val := range Values(s) {
		if f(val) {
			Add(t, val)
		}
	}
	return t
}

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

func Intersection[S Set[E], E comparable](s1, s2 S) Set[E] {
	s := Set[E]{}
	for val := range s1 {
		if Contains(s2, val) {
			Add(s, val)
		}
	}
	return s
}

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
