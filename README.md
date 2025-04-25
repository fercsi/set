# Golang Set implementation

This is a very simple implementation to use set-s.

```go
mySet := set.Set[int]{}

set.Add(mySet, 42)
if set.Contains(mySet, 137) {
    ...
}
set.Delete(23)

count = len(mySet)

for v := range mySet {
}
```
