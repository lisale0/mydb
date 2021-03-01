The operators follows an interface that contains Next and Execute.

```
type Operator interface {
	Next() bool
	Execute() Tuple
}
```

`Next()` checks if there is a valid tuple to be returned on the next call

`Execute()` actually returns the tuple.

TODO:
- Dynamically build the operators based on query