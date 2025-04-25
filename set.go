package set

import "iter"

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

// Add inserts the value `val` into the given set `s`.
//
// If the value already exists in the set, this operation has no effect.
//
// Parameters:
//   - s: The set to which the value will be added.
//   - val: The value to add to the set.
//
// Example:
//
//     s := make(Set[int])
//     Add(s, 42)
//
// Output:
//     The set will contain 42.
//
// Note:
//   This function modifies the set in-place.
func Add[E comparable](s Set[E], val E) {
    s[val] = struct{}{}
}

func Contains[E comparable](s Set[E], val E) bool {
    _, ok := s[val]
    return ok
}

func Delete[E comparable](s Set[E], val E) {
    delete(s, val)
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
