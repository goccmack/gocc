UserContext and $Context
========================

There are circumstances where you may need to provide stateful context
your parser actions. This can be done through the Parser.UserContext
field and then accessed in the actions with $Context, where it will be
exposed to you as an `interface{}`.

This example demonstrates using this to track filenames in concurrent
parsers.

See cmd/main.go for a minor example of using context, to execute it:

```
	go run ./cmd
```
