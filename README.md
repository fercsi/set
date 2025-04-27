# Golang Set implementation

This is a very simple implementation to use set-s.

## Installation

```bash
go get github.com/fercsi/set
```

## Usage

Example of how to import and use it:

```go
import "github.com/fercsi/aoc"


func main() {
    mySet := set.Set[int]{}

    set.Add(mySet, 42, 137)
    if set.Contains(mySet, 137) {
        ...
    }
    set.Delete(23, 75)

    count = len(mySet)

    for v := range mySet {
    }
}
```

## Features

- `type Set[E comparable] map[E]struct{}`
- `func Values[S Set[E], E comparable](s S) iter.Seq[E]`
- `func Add[S Set[E], E comparable](s S, vals ...E)`
- `func Contains[S Set[E], E comparable](s S, val E) bool`
- `func Delete[S Set[E], E comparable](s S, vals ...E)`
- `func DeleteSeq[S Set[E], E comparable](s S, seq iter.Seq[E])`
- `func Collect[E comparable](seq iter.Seq[E]) Set[E]`
- `func Map[S Set[E], E comparable, R comparable](s S, f func(E) R) Set[R]`
- `func Filter[S Set[E], E comparable](s S, f func(E) bool) Set[E]`
- `func Union[S Set[E], E comparable](s1, s2 S) Set[E]`
- `func Intersection[S Set[E], E comparable](s1, s2 S) Set[E]`
- `func (s Set[E]) String() string`

## Documentation

- [GoDoc reference](https://pkg.go.dev/github.com/fercsi/set)

## Testing

How to run tests:

```bash
go test
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
